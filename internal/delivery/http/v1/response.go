package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/http-server/pkg/logging"
)

type dataResponse struct {
	Data interface{} `json:"data"`
}

type idResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logging.GetLogger().Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
