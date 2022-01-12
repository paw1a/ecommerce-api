package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
		h.initProductsRoutes(v1)
		h.initReviewsRoutes(v1)
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
