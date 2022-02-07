package service

import (
	"context"
	"fmt"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartService struct {
	repo           repository.Carts
	productService Products
}

func (c *CartService) FindAll(ctx context.Context) ([]domain.Cart, error) {
	carts, err := c.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, cart := range carts {
		var totalPrice float64
		for _, cartItem := range cart.CartItems {
			product, err := c.productService.FindByID(ctx, cartItem.ProductID)
			if err != nil {
				return nil, fmt.Errorf("product with id %s no longer exist in stock", product.ID.Hex())
			}
			totalPrice += product.Price * float64(cartItem.Quantity)
		}
		carts[i].TotalPrice = totalPrice
	}

	return carts, nil
}

func (c *CartService) FindByID(ctx context.Context, cartID primitive.ObjectID) (domain.Cart, error) {
	cart, err := c.repo.FindByID(ctx, cartID)
	if err != nil {
		return domain.Cart{}, err
	}

	var totalPrice float64
	for _, cartItem := range cart.CartItems {
		product, err := c.productService.FindByID(ctx, cartItem.ProductID)
		if err != nil {
			return domain.Cart{}, fmt.Errorf("product with id %s no longer exist in stock", product.ID.Hex())
		}
		totalPrice += product.Price * float64(cartItem.Quantity)
	}
	cart.TotalPrice = totalPrice

	return cart, nil
}

func (c *CartService) FindCartItems(ctx context.Context, cartID primitive.ObjectID) ([]domain.CartItem, error) {
	return c.repo.FindCartItems(ctx, cartID)
}

func (c *CartService) AddCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error) {
	return c.repo.AddCartItem(ctx, cartItem, cartID)
}

func (c *CartService) UpdateCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error) {
	return c.repo.UpdateCartItem(ctx, cartItem, cartID)
}

func (c *CartService) DeleteCartItem(ctx context.Context, productID primitive.ObjectID, cartID primitive.ObjectID) error {
	return c.repo.DeleteCartItem(ctx, productID, cartID)
}

func (c *CartService) ClearCart(ctx context.Context, cartID primitive.ObjectID) error {
	return c.repo.ClearCart(ctx, cartID)
}

func (c *CartService) Create(ctx context.Context, cartDTO dto.CreateCartDTO) (domain.Cart, error) {
	return c.repo.Create(ctx, domain.Cart{
		ExpireAt:  cartDTO.ExpireAt,
		CartItems: cartDTO.CartItems,
	})
}

func (c *CartService) Update(ctx context.Context, cartDTO dto.UpdateCartDTO, cartID primitive.ObjectID) (domain.Cart, error) {
	return c.repo.Update(ctx, dto.UpdateCartInput{
		ExpireAt:  cartDTO.ExpireAt,
		CartItems: cartDTO.CartItems,
	}, cartID)
}

func (c *CartService) Delete(ctx context.Context, cartID primitive.ObjectID) error {
	return c.repo.Delete(ctx, cartID)
}

func NewCartsService(repo repository.Carts, productsService Products) *CartService {
	return &CartService{
		repo:           repo,
		productService: productsService,
	}
}
