package service

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentService struct {
}

func (p *PaymentService) GetPaymentUrl(order domain.Order) (string, error) {
	sessionParamsList := make([]*stripe.CheckoutSessionLineItemParams, len(order.OrderItems))
	for i, orderItem := range order.OrderItems {
		productPrice := p.GetProductPrice(orderItem.ProductID)
		sessionParamsList[i] = &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(productPrice.ID),
			Quantity: stripe.Int64(orderItem.Quantity),
		}
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:         sessionParamsList,
		SuccessURL:        stripe.String("http://localhost:8080/"),
		CancelURL:         stripe.String("http://localhost:8080/"),
		Mode:              stripe.String(string(stripe.CheckoutSessionModePayment)),
		ClientReferenceID: stripe.String(order.ID.Hex()),
	}
	checkoutSession, err := session.New(params)

	if err != nil {
		return "", err
	}

	return checkoutSession.URL, nil
}

func (p *PaymentService) GetProductPrice(productID primitive.ObjectID) *stripe.Price {
	params := &stripe.PriceListParams{Product: stripe.String(productID.Hex())}
	iterator := price.List(params)
	iterator.Next()
	return iterator.Price()
}

func (p *PaymentService) CreateProduct(domainProduct domain.Product) error {
	productParams := &stripe.ProductParams{
		ID:          stripe.String(domainProduct.ID.Hex()),
		Name:        stripe.String(domainProduct.Name),
		Description: stripe.String(domainProduct.Description),
	}
	stripeProduct, err := product.New(productParams)
	if err != nil {
		return err
	}

	priceParams := &stripe.PriceParams{
		Currency:          stripe.String(string(stripe.CurrencyRUB)),
		Product:           stripe.String(stripeProduct.ID),
		UnitAmountDecimal: stripe.Float64(domainProduct.Price * 100),
	}

	_, err = price.New(priceParams)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentService) UpdateProduct(domainProduct domain.Product) error {
	params := &stripe.PriceListParams{Product: stripe.String(domainProduct.ID.Hex())}
	iterator := price.List(params)
	iterator.Next()
	productPrice := iterator.Price()
	_, err := price.Update(productPrice.ID, &stripe.PriceParams{
		UnitAmountDecimal: stripe.Float64(domainProduct.Price * 100),
	})
	if err != nil {
		return err
	}

	_, err = product.Update(domainProduct.ID.Hex(), &stripe.ProductParams{
		Name:        stripe.String(domainProduct.Name),
		Description: stripe.String(domainProduct.Description),
	})

	return err
}

func (p *PaymentService) DeleteProduct(productID primitive.ObjectID) error {
	_, err := product.Del(productID.Hex(), nil)
	return err
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}
