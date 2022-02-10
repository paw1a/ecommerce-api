package service

import (
	"context"
	"fmt"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrdersService struct {
	repo           repository.Orders
	productService Products
	cartService    Carts
}

func (p *OrdersService) FindAll(ctx context.Context) ([]domain.Order, error) {
	orders, err := p.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, order := range orders {
		var totalPrice float64
		for _, orderItem := range order.OrderItems {
			product, err := p.productService.FindByID(ctx, orderItem.ProductID)
			if err != nil {
				return nil, err
			}
			totalPrice += product.Price * float64(orderItem.Quantity)
		}
		orders[i].TotalPrice = totalPrice
	}

	return orders, nil
}

func (p *OrdersService) FindByID(ctx context.Context, orderID primitive.ObjectID) (domain.Order, error) {
	order, err := p.repo.FindByID(ctx, orderID)
	if err != nil {
		return domain.Order{}, err
	}

	var totalPrice float64
	for _, orderItem := range order.OrderItems {
		product, err := p.productService.FindByID(ctx, orderItem.ProductID)
		if err != nil {
			return domain.Order{}, err
		}
		totalPrice += product.Price * float64(orderItem.Quantity)
	}
	order.TotalPrice = totalPrice

	return order, err
}

func (p *OrdersService) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Order, error) {
	orders, err := p.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i, order := range orders {
		var totalPrice float64
		for _, orderItem := range order.OrderItems {
			product, err := p.productService.FindByID(ctx, orderItem.ProductID)
			if err != nil {
				return nil, err
			}
			totalPrice += product.Price * float64(orderItem.Quantity)
		}
		orders[i].TotalPrice = totalPrice
	}

	return orders, nil
}

func (p *OrdersService) Create(ctx context.Context, orderDTO dto.CreateOrderDTO) (domain.Order, error) {
	var totalPrice float64
	for _, orderItem := range orderDTO.OrderItems {
		product, err := p.productService.FindByID(ctx, orderItem.ProductID)
		if err != nil {
			return domain.Order{}, fmt.Errorf("product with id %s no longer exist in stock", product.ID.Hex())
		}
		totalPrice += product.Price * float64(orderItem.Quantity)
	}

	return p.repo.Create(ctx, domain.Order{
		OrderID:     uuid.NewV4().String(),
		CreatedAt:   time.Now(),
		OrderItems:  orderDTO.OrderItems,
		ContactInfo: orderDTO.ContactInfo,
		TotalPrice:  totalPrice,
		UserID:      orderDTO.UserID,
		Status:      "reserved",
	})
}

func (p *OrdersService) Update(ctx context.Context, orderDTO dto.UpdateOrderDTO, orderID primitive.ObjectID) (domain.Order, error) {
	return p.repo.Update(ctx, dto.UpdateOrderInput{
		DeliveredAt: orderDTO.DeliveredAt,
		Status:      orderDTO.Status,
	}, orderID)
}

func (p *OrdersService) Delete(ctx context.Context, orderID primitive.ObjectID) error {
	return p.repo.Delete(ctx, orderID)
}

func NewOrdersService(repo repository.Orders, productService Products, cartService Carts) *OrdersService {
	return &OrdersService{
		repo:           repo,
		productService: productService,
		cartService:    cartService,
	}
}
