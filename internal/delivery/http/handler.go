package http

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/config"
	v1 "github.com/paw1a/ecommerce-api/internal/delivery/http/v1"
	"github.com/paw1a/ecommerce-api/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	// Init gin handler
	router := gin.Default()

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
