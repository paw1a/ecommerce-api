package service_test

import (
	"context"
	"errors"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/repository/mocks"
	"github.com/paw1a/ecommerce-api/internal/service"
	mocks_service "github.com/paw1a/ecommerce-api/internal/service/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/stretchr/testify.v1/require"
	"testing"
)

var testProducts = []domain.Product{
	domain.Product{
		ID:    primitive.NewObjectID(),
		Name:  "First product",
		Price: 1200,
	},
	domain.Product{
		ID:    primitive.NewObjectID(),
		Name:  "Second product",
		Price: 1500,
	},
}

var testCart = domain.Cart{
	ID: primitive.NewObjectID(),
	CartItems: []domain.CartItem{
		{
			ProductID: testProducts[0].ID,
			Quantity:  3,
		},
		{
			ProductID: testProducts[1].ID,
			Quantity:  2,
		},
	},
}

var unknownID = primitive.NewObjectID()

func TestCartService_FindByID(t *testing.T) {
	testTable := []struct {
		name        string
		inputCart   domain.Cart
		cartFunc    func(cartRepo *mocks_repository.Carts)
		productFunc func(productService *mocks_service.Products)
		price       float64
		err         error
	}{
		{
			name:      "valid and found",
			inputCart: testCart,
			cartFunc: func(cartRepo *mocks_repository.Carts) {
				cartRepo.On("FindByID", context.Background(), testCart.ID).
					Return(testCart, nil)
			},
			productFunc: func(productService *mocks_service.Products) {
				productService.On("FindByID", context.Background(), testProducts[0].ID).
					Return(testProducts[0], nil)
				productService.On("FindByID", context.Background(), testProducts[1].ID).
					Return(testProducts[1], nil)
			},
			price: 6600,
			err:   nil,
		},
		{
			name: "cart not found",
			inputCart: domain.Cart{
				ID: unknownID,
			},
			cartFunc: func(cartRepo *mocks_repository.Carts) {
				cartRepo.On("FindByID", context.Background(), unknownID).
					Return(domain.Cart{}, errors.New("cart not found"))
			},
			productFunc: func(productService *mocks_service.Products) {

			},
			err: errors.New("cart not found error"),
		},
		{
			name:      "no products in cart",
			inputCart: testCart,
			cartFunc: func(cartRepo *mocks_repository.Carts) {
				cartRepo.On("FindByID", context.Background(), testCart.ID).
					Return(domain.Cart{
						ID:        testCart.ID,
						CartItems: []domain.CartItem{},
					}, nil)
			},
			productFunc: func(productService *mocks_service.Products) {

			},
			price: 0,
			err:   nil,
		},
		{
			name:      "one of the products not found",
			inputCart: testCart,
			cartFunc: func(cartRepo *mocks_repository.Carts) {
				cartRepo.On("FindByID", context.Background(), testCart.ID).
					Return(domain.Cart{
						ID: testCart.ID,
						CartItems: []domain.CartItem{
							{
								ProductID: testProducts[0].ID,
								Quantity:  2,
							},
							{
								ProductID: unknownID,
								Quantity:  3,
							},
						},
					}, nil)
			},
			productFunc: func(productService *mocks_service.Products) {
				productService.On("FindByID", context.Background(), testProducts[0].ID).
					Return(testProducts[0], nil)
				productService.On("FindByID", context.Background(), unknownID).
					Return(domain.Product{}, errors.New("product with this id not found"))
			},
			err: errors.New("one of the products is out of stock"),
		},
	}

	for _, test := range testTable {
		t.Logf("Test: %s", test.name)

		cartRepo := mocks_repository.NewCarts(t)
		productsService := mocks_service.NewProducts(t)

		cartService := service.NewCartsService(cartRepo, productsService)

		test.productFunc(productsService)
		test.cartFunc(cartRepo)

		cart, err := cartService.FindByID(context.Background(), test.inputCart.ID)

		require.Equal(t, test.price, cart.TotalPrice)
		if test.err != nil {
			require.Error(t, err)
		}
	}
}
