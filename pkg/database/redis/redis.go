package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/paw1a/ecommerce-api/internal/config"
)

func NewClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.URI,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
