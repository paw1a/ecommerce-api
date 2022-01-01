package app

import (
	"context"
	"fmt"
	"github.com/paw1a/http-server/internal/config"
	delivery "github.com/paw1a/http-server/internal/delivery/http"
	"github.com/paw1a/http-server/internal/repository"
	"github.com/paw1a/http-server/internal/service"
	"github.com/paw1a/http-server/pkg/database/mongodb"
	"github.com/paw1a/http-server/pkg/logging"
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

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := delivery.NewHandler(services)

	server := &http.Server{
		Handler:      handlers.Init(cfg),
		Addr:         fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
