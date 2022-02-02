package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{
		repo: repo,
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
	return u.repo.Create(ctx, domain.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	})
}

func (u *UsersService) Update(ctx context.Context, userDTO dto.UpdateUserDTO, userID primitive.ObjectID) (domain.User, error) {
	return u.repo.Update(ctx, dto.UpdateUserInput{
		Name:   userDTO.Name,
		CartID: userDTO.CartID,
	}, userID)
}

func (u *UsersService) Delete(ctx context.Context, userID primitive.ObjectID) error {
	return u.repo.Delete(ctx, userID)
}
