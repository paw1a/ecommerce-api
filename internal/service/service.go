package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
}

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, product dto.CreateProductDTO) (domain.Product, error)
	Update(ctx context.Context, productDTO dto.UpdateProductDTO) (domain.Product, error)
	Delete(ctx context.Context, productID primitive.ObjectID) error
}

type Reviews interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error)
	Create(ctx context.Context, review dto.CreateReviewDTO) (domain.Review, error)
	Delete(ctx context.Context, reviewID primitive.ObjectID) error
}

type Services struct {
	Users    Users
	Products Products
	Reviews  Reviews
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	productsService := NewProductsService(deps.Repos.Products)
	reviewsService := NewReviewsService(deps.Repos.Reviews)

	return &Services{
		Users:    usersService,
		Products: productsService,
		Reviews:  reviewsService,
	}
}
