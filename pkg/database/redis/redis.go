package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/paw1a/ecommerce-api/internal/config"
	log "github.com/sirupsen/logrus"
)

func NewClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.URI,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error("failed to ping redis")
		return nil, err
	}

	return client, nil
}
