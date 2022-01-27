package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/pkg/auth"
	"net/http"
	"strings"
)

func extractAuthToken(context *gin.Context) (string, error) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

func (h *Handler) refreshToken(context *gin.Context) {
	var input auth.RefreshInput
	err := context.BindJSON(&input)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid request body")
		return
	}

	authDetails, err := h.tokenProvider.Refresh(auth.RefreshInput{
		RefreshToken: input.RefreshToken,
		Fingerprint:  input.Fingerprint,
	})

	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}
	successResponse(context, authDetails)
}

func (h *Handler) verifyToken(context *gin.Context, idName string) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	id, ok := tokenClaims[idName]
	if !ok {
		errorResponse(context, http.StatusForbidden, "this endpoint is forbidden")
		return
	}

	context.Set(idName, id)
}
