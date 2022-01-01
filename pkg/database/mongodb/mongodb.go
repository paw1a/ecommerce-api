package mongodb

import (
	"context"
	"github.com/paw1a/http-server/internal/config"
	"github.com/paw1a/http-server/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, config *config.Config, logger *logging.Logger) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(config.DB.URI)
	if config.DB.Username != "" && config.DB.Password != "" {
		opts.SetAuth(options.Credential{
			Username: config.DB.Username, Password: config.DB.Password,
		})
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		logger.Errorf("Failed to connect to mongo")
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Errorf("Failed to ping mongo")
		return nil, err
	}

	return client, nil
}
