package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/pkg/auth"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	log.Infof("%v", context.Request.Header)

	err := context.BindJSON(&input)
	if err != nil {
		badRequestResponse(context, "can't parse request body", err)
		return
	}

	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		badRequestResponse(context, "refresh cookie not found", err)
		return
	}

	input.RefreshToken = refreshCookie

	authDetails, err := h.tokenProvider.Refresh(auth.RefreshInput{
		RefreshToken: input.RefreshToken,
		Fingerprint:  input.Fingerprint,
	})

	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken,
		int(h.config.JWT.RefreshTokenTime), "/", h.config.Listen.Host, false, false)

	successResponse(context, authDetails.AccessToken)
}

func (h *Handler) verifyToken(context *gin.Context, idName string) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	id, ok := tokenClaims[idName]
	if !ok {
		errorResponse(context, http.StatusForbidden, "this endpoint is forbidden")
		return
	}

	context.Set(idName, id)
}

func (h *Handler) extractIdFromAuthHeader(context *gin.Context, idName string) (primitive.ObjectID, error) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	idHex, ok := tokenClaims[idName]
	if !ok {
		return primitive.ObjectID{}, fmt.Errorf("failed to extract %s from auth header", idName)
	}

	id, err := primitive.ObjectIDFromHex(idHex.(string))
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("failed to convert %s to objectId", idHex)
	}

	return id, nil
}
