package payment

import (
	"github.com/paw1a/ecommerce-api/internal/config"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhookendpoint"
)

func InitStripeClient(cfg *config.Config) error {
	stripe.Key = cfg.Stripe.Key

	params := &stripe.WebhookEndpointListParams{}
	i := webhookendpoint.List(params)
	for i.Next() {
		webhook := i.WebhookEndpoint()
		if webhook.URL == cfg.Stripe.WebhookUrl {
			cfg.Stripe.WebhookSecret = webhook.Secret
			return nil
		}
	}

	webhookParams := &stripe.WebhookEndpointParams{
		EnabledEvents: []*string{
			stripe.String("checkout.session.completed"),
		},
		URL: stripe.String(cfg.Stripe.WebhookUrl),
	}
	webhook, err := webhookendpoint.New(webhookParams)

	cfg.Stripe.WebhookSecret = webhook.Secret
	return err
}
