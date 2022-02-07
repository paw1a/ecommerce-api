package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error)
	FindByCredentials(ctx context.Context, email string, password string) (domain.User, error)
	FindUserInfo(ctx context.Context, userID primitive.ObjectID) (domain.UserInfo, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, userInput dto.UpdateUserInput,
		userID primitive.ObjectID) (domain.User, error)
	Delete(ctx context.Context, userID primitive.ObjectID) error
}

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, productInput dto.UpdateProductInput,
		productID primitive.ObjectID) (domain.Product, error)
	Delete(ctx context.Context, productID primitive.ObjectID) error
}

type Reviews interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Review, error)
	FindByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Review, error)
	Create(ctx context.Context, review domain.Review) (domain.Review, error)
	Delete(ctx context.Context, reviewID primitive.ObjectID) error
	DeleteByProductID(ctx context.Context, productID primitive.ObjectID) error
}

type Admins interface {
	FindByCredentials(ctx context.Context, email string, password string) (domain.Admin, error)
}

type Carts interface {
	FindAll(ctx context.Context) ([]domain.Cart, error)
	FindByID(ctx context.Context, cartID primitive.ObjectID) (domain.Cart, error)
	FindCartItems(ctx context.Context, cartID primitive.ObjectID) ([]domain.CartItem, error)
	AddCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error)
	DeleteCartItem(ctx context.Context, productID primitive.ObjectID, cartID primitive.ObjectID) error
	ClearCart(ctx context.Context, cartID primitive.ObjectID) error
	Create(ctx context.Context, cart domain.Cart) (domain.Cart, error)
	Update(ctx context.Context, cartInput dto.UpdateCartInput,
		cartID primitive.ObjectID) (domain.Cart, error)
	Delete(ctx context.Context, cartID primitive.ObjectID) error
}

type Orders interface {
	FindAll(ctx context.Context) ([]domain.Order, error)
	FindByID(ctx context.Context, orderID primitive.ObjectID) (domain.Order, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Order, error)
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, orderInput dto.UpdateOrderInput,
		orderID primitive.ObjectID) (domain.Order, error)
	Delete(ctx context.Context, orderID primitive.ObjectID) error
}

type Repositories struct {
	Users    Users
	Products Products
	Reviews  Reviews
	Admins   Admins
	Carts    Carts
	Orders   Orders
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db),
		Products: NewProductsRepo(db),
		Reviews:  NewReviewsRepo(db),
		Admins:   NewAdminsRepo(db),
		Carts:    NewCartsRepo(db),
		Orders:   NewOrdersRepo(db),
	}
}
