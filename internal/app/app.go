package app

import (
	"context"
	"fmt"
	"github.com/paw1a/ecommerce-api/internal/config"
	delivery "github.com/paw1a/ecommerce-api/internal/delivery/http"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"github.com/paw1a/ecommerce-api/internal/service"
	"github.com/paw1a/ecommerce-api/pkg/auth"
	"github.com/paw1a/ecommerce-api/pkg/database/mongodb"
	"github.com/paw1a/ecommerce-api/pkg/database/redis"
	"github.com/paw1a/ecommerce-api/pkg/logging"
	"log"
	"net/http"
	"time"
)

func Run(configPath string) {
	logger := logging.GetLogger()
	logger.Info("Logger is created")

	cfg := config.GetConfig(configPath)
	logger.Info("Config is created")

	mongoClient, err := mongodb.NewClient(context.Background(), cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Mongo is connected")
	db := mongoClient.Database(cfg.DB.Database)

	redisClient, err := redis.NewClient(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Redis is connected")

	tokenProvider := auth.NewTokenProvider(cfg, redisClient)

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	handlers := delivery.NewHandler(services, tokenProvider)

	server := &http.Server{
		Handler:      handlers.Init(cfg),
		Addr:         fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
