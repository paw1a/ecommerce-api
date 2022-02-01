package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
		for _, cartItem := range cart.Products {
			product, err := c.productService.FindByID(ctx, cartItem.ProductID)
			if err != nil {
				return nil, err
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
	for _, cartItem := range cart.Products {
		product, err := c.productService.FindByID(ctx, cartItem.ProductID)
		if err != nil {
			return domain.Cart{}, err
		}
		totalPrice += product.Price * float64(cartItem.Quantity)
	}
	cart.TotalPrice = totalPrice

	return cart, nil
}

func (c *CartService) Create(ctx context.Context, cartDTO dto.CreateCartDTO) (domain.Cart, error) {
	return c.repo.Create(ctx, domain.Cart{
		CreatedAt: time.Now(),
		Products:  cartDTO.Products,
	})
}

func (c *CartService) Update(ctx context.Context, cartDTO dto.UpdateCartDTO, cartID primitive.ObjectID) (domain.Cart, error) {
	return c.repo.Update(ctx, dto.UpdateCartInput{Products: cartDTO.Products}, cartID)
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
