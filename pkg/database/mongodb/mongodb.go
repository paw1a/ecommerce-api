package mongodb

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, uri string, dbUsername string, dbPassword string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri)
	if dbUsername != "" && dbPassword != "" {
		opts.SetAuth(options.Credential{
			Username: dbUsername, Password: dbPassword,
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
