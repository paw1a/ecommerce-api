package v1

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/pkg/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// UserSignIn godoc
// @Summary  User sign-in
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Param    user  body      dto.SignInDTO  true  "user credentials"
// @Success  200   {object}  auth.AuthDetails
// @Failure  400   {object}  failure
// @Failure  401   {object}  failure
// @Failure  404   {object}  failure
// @Failure  500   {object}  failure
// @Router   /users/auth/sign-in [post]
func (h *Handler) userSignIn(context *gin.Context) {
	var signInDTO dto.SignInDTO
	err := context.BindJSON(&signInDTO)
	if err != nil {
		badRequestResponse(context, "invalid sign in credentials", err)
		return
	}

	user, err := h.services.Users.FindByCredentials(context, signInDTO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			unauthorizedResponse(context, "invalid user email or password")
		} else {
			internalErrorResponse(context, err)
		}
		return
	}

	if err != nil {

	}

	userClaims := jwt.MapClaims{"userID": user.ID}
	authDetails, err := h.tokenProvider.CreateJWTSession(auth.CreateSessionInput{
		Fingerprint: signInDTO.Fingerprint,
		Claims:      userClaims,
	})

	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken,
		int(h.config.JWT.RefreshTokenTime), "/", h.config.Listen.Host, false, false)

	successResponse(context, authDetails.AccessToken)
}

// UserSignUp godoc
// @Summary  User sign-up
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Param    user  body      dto.SignUpDTO  true  "user data"
// @Success  200   {object}  domain.UserInfo
// @Failure  400   {object}  failure
// @Failure  401   {object}  failure
// @Failure  404   {object}  failure
// @Failure  500   {object}  failure
// @Router   /users/auth/sign-up [post]
func (h *Handler) userSignUp(context *gin.Context) {
	var signUpDTO dto.SignUpDTO
	err := context.BindJSON(&signUpDTO)
	if err != nil {
		badRequestResponse(context, "invalid sign up data", err)
		return
	}

	user, err := h.services.Users.Create(context, dto.CreateUserDTO{
		Name:     signUpDTO.Name,
		Email:    signUpDTO.Email,
		Password: signUpDTO.Password,
		CartID:   signUpDTO.CartID,
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			badRequestResponse(context,
				fmt.Sprintf("user with email %s already exists", signUpDTO.Email), err)
		} else {
			internalErrorResponse(context, err)
		}
		return
	}

	createdResponse(context, domain.UserInfo{
		Name:  user.Name,
		Email: user.Email,
	})
}

// UserRefresh godoc
// @Summary  User refresh token
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Param    refreshInput  body      auth.RefreshInput  true  "user refresh data"
// @Success  200           {object}  auth.AuthDetails
// @Failure  400           {object}  failure
// @Failure  401           {object}  failure
// @Failure  404           {object}  failure
// @Failure  500           {object}  failure
// @Router   /users/auth/refresh [post]
func (h *Handler) userRefresh(context *gin.Context) {
	h.refreshToken(context)
}

func (h *Handler) verifyUser(context *gin.Context) {
	h.verifyToken(context, "userID")
}
