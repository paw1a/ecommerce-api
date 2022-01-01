package service

import (
	"context"
	"github.com/paw1a/http-server/internal/domain"
	"github.com/paw1a/http-server/internal/domain/dto"
	"github.com/paw1a/http-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
}

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, Product dto.CreateProductDTO) (domain.Product, error)
	Update(ctx context.Context, productDTO dto.UpdateProductDTO) (domain.Product, error)
	Delete(ctx context.Context, productID primitive.ObjectID) error
}

type Services struct {
	Users    Users
	Products Products
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	productsService := NewProductsService(deps.Repos.Products)

	return &Services{
		Users:    usersService,
		Products: productsService,
	}
}
