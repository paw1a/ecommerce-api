package repository

import (
	"context"
	"github.com/paw1a/http-server/internal/domain"
	"github.com/paw1a/http-server/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
}

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, productInput dto.UpdateProductInput) (domain.Product, error)
	Delete(ctx context.Context, productID primitive.ObjectID) error
}

type Repositories struct {
	Users    Users
	Products Products
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db),
		Products: NewProductsRepo(db),
	}
}
