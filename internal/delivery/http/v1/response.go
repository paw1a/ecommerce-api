package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type success struct {
	Data interface{} `json:"data"`
}

type failure struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"invalid request body"`
}

func successResponse(c *gin.Context, data interface{}) {
	log.Infof("Response OK: %v", data)
	c.JSON(http.StatusOK, success{Data: data})
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, failure{
		Code:    statusCode,
		Message: message,
	})
}
