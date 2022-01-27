package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error)
	FindByCredentials(ctx context.Context, signInDTO dto.SignInDTO) (domain.User, error)
	FindUserInfo(ctx context.Context, userID primitive.ObjectID) (domain.UserInfo, error)
	Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error)
	Update(ctx context.Context, userDTO dto.UpdateUserDTO,
		userID primitive.ObjectID) (domain.User, error)
	Delete(ctx context.Context, userID primitive.ObjectID) error
}

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, productDTO dto.CreateProductDTO) (domain.Product, error)
	Update(ctx context.Context, productDTO dto.UpdateProductDTO,
		productID primitive.ObjectID) (domain.Product, error)
	Delete(ctx context.Context, productID primitive.ObjectID) error
}

type Reviews interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Review, error)
	FindByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Review, error)
	Create(ctx context.Context, review dto.CreateReviewInput) (domain.Review, error)
	Delete(ctx context.Context, reviewID primitive.ObjectID) error
}

type Admins interface {
	FindByCredentials(ctx context.Context, signInDTO dto.SignInDTO) (domain.Admin, error)
}

type Services struct {
	Users    Users
	Products Products
	Reviews  Reviews
	Admins   Admins
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	productsService := NewProductsService(deps.Repos.Products)
	reviewsService := NewReviewsService(deps.Repos.Reviews)
	adminsService := NewAdminsService(deps.Repos.Admins)

	return &Services{
		Users:    usersService,
		Products: productsService,
		Reviews:  reviewsService,
		Admins:   adminsService,
	}
}
