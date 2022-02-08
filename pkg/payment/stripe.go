package payment

import (
	"github.com/paw1a/ecommerce-api/internal/config"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhookendpoint"
)

func InitStripeClient(cfg *config.Config) error {
	stripe.Key = cfg.Stripe.Key

	webhookParams := &stripe.WebhookEndpointParams{
		EnabledEvents: []*string{
			stripe.String("charge.failed"),
			stripe.String("charge.succeeded"),
		},
		URL: stripe.String("http://localhost:8080/v1/api/payment/webhook"),
	}
	_, err := webhookendpoint.New(webhookParams)
	return err
}
