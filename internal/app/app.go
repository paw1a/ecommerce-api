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
	_ "github.com/paw1a/ecommerce-api/pkg/logging"
	"github.com/paw1a/ecommerce-api/pkg/payment"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Run(configPath string) {
	log.Info("application startup...")
	log.Info("logger initialized")

	cfg := config.GetConfig(configPath)
	log.Info("config created")

	mongoClient, err := mongodb.NewClient(context.Background(),
		cfg.DB.URI, cfg.DB.Username, cfg.DB.Password)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("mongo is connected")
	db := mongoClient.Database(cfg.DB.Database)

	redisClient, err := redis.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("redis is connected")

	tokenProvider := auth.NewTokenProvider(cfg, redisClient)
	log.Info("token provider initialized")

	err = payment.InitStripeClient(cfg)
	if err != nil {
		log.Error(err)
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:       repos,
		RedisClient: redisClient,
	})
	handlers := delivery.NewHandler(services, tokenProvider, cfg)
	log.Info("services, repositories and handlers initialized")

	server := &http.Server{
		Handler:      handlers.Init(),
		Addr:         fmt.Sprintf(":%s", cfg.Listen.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("server started on port %s", cfg.Listen.Port)

	log.Fatal(server.ListenAndServe())
}
