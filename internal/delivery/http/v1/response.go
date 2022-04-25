package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type success struct {
	Data interface{} `json:"data"`
}

type failure struct {
	Code    int    `json:"code" example:"403"`
	Message string `json:"message" example:"resource is forbidden"`
}

type notFoundError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"resource not found"`
}

type badRequestError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"invalid request body"`
}

type internalError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"something went wrong on the server"`
	Error   error  `json:"-"`
}

type unauthorizedError struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"user with id=123 is unauthorized"`
}

func successResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, success{Data: data})
}

func createdResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusCreated, success{Data: data})
}

func internalErrorResponse(context *gin.Context, err error) {
	response := internalError{
		Code:    http.StatusInternalServerError,
		Message: "500 Internal Server Error, contact us to fix it",
		Error:   err,
	}
	log.Errorf("Something went wrong: %v", response.Error)
	context.AbortWithStatusJSON(response.Code, response.Error)
}

func badRequestResponse(context *gin.Context, message string, err error) {
	response := badRequestError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
	log.Warnf("User bad request error: %v", err)
	context.AbortWithStatusJSON(response.Code, message)
}

func unauthorizedResponse(context *gin.Context, message string) {

}

func notFoundOrInternalErrorResponse(context *gin.Context, notFoundMessage string, err error) {
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Warn("Not Found error: ", notFoundMessage)
		context.AbortWithStatusJSON(http.StatusNotFound, notFoundMessage)
	} else {
		internalErrorResponse(context, err)
	}
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, failure{
		Code:    statusCode,
		Message: message,
	})
}
