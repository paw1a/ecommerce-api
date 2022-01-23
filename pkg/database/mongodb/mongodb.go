package mongodb

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, config *config.Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(config.DB.URI)
	if config.DB.Username != "" && config.DB.Password != "" {
		opts.SetAuth(options.Credential{
			Username: config.DB.Username, Password: config.DB.Password,
		})
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Error("failed to create mongo client")
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Errorf("failed to connect to mongo")
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Errorf("failed to ping mongo")
		return nil, err
	}

	return client, nil
}
