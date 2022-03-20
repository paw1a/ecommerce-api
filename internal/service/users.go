package service

import (
	"context"
	"fmt"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UsersService struct {
	repo        repository.Users
	cartService Carts
}

func NewUsersService(repo repository.Users, cartService Carts) *UsersService {
	return &UsersService{
		repo:        repo,
		cartService: cartService,
	}
}

func (u *UsersService) FindAll(ctx context.Context) ([]domain.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *UsersService) FindByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error) {
	return u.repo.FindByID(ctx, userID)
}

func (u *UsersService) FindByCredentials(ctx context.Context, signInDTO dto.SignInDTO) (domain.User, error) {
	return u.repo.FindByCredentials(ctx, signInDTO.Email, signInDTO.Password)
}

func (u *UsersService) FindUserInfo(ctx context.Context, userID primitive.ObjectID) (domain.UserInfo, error) {
	return u.repo.FindUserInfo(ctx, userID)
}

func (u UsersService) Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error) {
	var cartID primitive.ObjectID
	expireTime := time.Now().Add(30 * time.Hour * 24)

	if userDTO.CartID == primitive.NilObjectID {
		cart, err := u.cartService.Create(ctx, dto.CreateCartDTO{
			ExpireAt: expireTime,
		})
		if err != nil {
			return domain.User{}, err
		}
		cartID = cart.ID
	} else {
		_, err := u.cartService.Update(ctx, dto.UpdateCartDTO{ExpireAt: &expireTime}, userDTO.CartID)
		if err != nil {
			return domain.User{}, err
		}
		cartID = userDTO.CartID
	}

	return u.repo.Create(ctx, domain.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		CartID:   cartID,
	})
}

func (u *UsersService) Update(ctx context.Context, userDTO dto.UpdateUserDTO, userID primitive.ObjectID) (domain.User, error) {
	return u.repo.Update(ctx, dto.UpdateUserInput{
		Name:   userDTO.Name,
		CartID: userDTO.CartID,
	}, userID)
}

func (u *UsersService) Delete(ctx context.Context, userID primitive.ObjectID) error {
	user, err := u.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("not found user with id: %s", userID)
	}
	err = u.repo.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return u.cartService.Delete(ctx, user.CartID)
}
