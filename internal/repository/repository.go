package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
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

type Reviews interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Review, error)
	FindByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Review, error)
	Create(ctx context.Context, review domain.Review) (domain.Review, error)
	Delete(ctx context.Context, reviewID primitive.ObjectID) error
}

type Admins interface {
	FindByCredentials(ctx context.Context, email string, password string) (domain.Admin, error)
}

type Repositories struct {
	Users    Users
	Products Products
	Reviews  Reviews
	Admins   Admins
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db),
		Products: NewProductsRepo(db),
		Reviews:  NewReviewsRepo(db),
	}
}
