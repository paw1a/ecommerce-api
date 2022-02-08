package service

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentlink"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentService struct {
	orderService   Orders
	productService Products
}

func (p *PaymentService) GetPaymentLink(ctx context.Context, orderID primitive.ObjectID) (string, error) {
	order, err := p.orderService.FindByID(ctx, orderID)
	if err != nil {
		return "", err
	}

	linkParamsList := make([]*stripe.PaymentLinkLineItemParams, len(order.OrderItems))
	for i, orderItem := range order.OrderItems {
		orderProduct, err := p.productService.FindByID(ctx, orderItem.ProductID)
		if err != nil {
			return "", fmt.Errorf("no products with id: %s", orderItem.ProductID.Hex())
		}

		productParams := &stripe.ProductParams{
			ID:          stripe.String(orderProduct.ID.Hex()),
			Name:        stripe.String(orderProduct.Name),
			Description: stripe.String(orderProduct.Description),
		}
		stripeProduct, err := product.New(productParams)
		if err != nil {
			return "", err
		}

		priceParams := &stripe.PriceParams{
			Currency:          stripe.String(string(stripe.CurrencyRUB)),
			Product:           stripe.String(stripeProduct.ID),
			UnitAmountDecimal: stripe.Float64(orderProduct.Price * float64(orderItem.Quantity)),
		}
		stripePrice, err := price.New(priceParams)
		if err != nil {
			return "", err
		}

		linkParamsList[i] = &stripe.PaymentLinkLineItemParams{
			Price:    stripe.String(stripePrice.ID),
			Quantity: stripe.Int64(orderItem.Quantity),
		}
	}

	params := &stripe.PaymentLinkParams{
		LineItems: linkParamsList,
	}
	paymentLink, err := paymentlink.New(params)
	if err != nil {
		return "", err
	}

	return paymentLink.URL, nil
}

func NewPaymentService(orderService Orders, productService Products) *PaymentService {
	return &PaymentService{
		orderService:   orderService,
		productService: productService,
	}
}
