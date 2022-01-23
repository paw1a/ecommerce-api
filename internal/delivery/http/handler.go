package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/paw1a/ecommerce-api/internal/delivery/http/v1"
	"github.com/paw1a/ecommerce-api/internal/service"
	"github.com/paw1a/ecommerce-api/pkg/auth"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/paw1a/ecommerce-api/docs"
	"net/http"
)

type Handler struct {
	services      *service.Services
	tokenProvider auth.TokenProvider
}

func NewHandler(services *service.Services, tokenProvider auth.TokenProvider) *Handler {
	return &Handler{
		services:      services,
		tokenProvider: tokenProvider,
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenProvider)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
