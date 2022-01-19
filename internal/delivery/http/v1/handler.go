package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/service"
	"github.com/paw1a/ecommerce-api/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAdminsRoutes(v1)
	}
}

func parseIdFromPath(c *gin.Context, paramName string) (primitive.ObjectID, error) {
	idParam := c.Param(paramName)
	if idParam == "" {
		return primitive.ObjectID{}, errors.New("empty id param")
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id param")
	}

	return id, nil
}
