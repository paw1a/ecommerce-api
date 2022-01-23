package v1

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/pkg/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/auth/sign-in", h.adminSignIn)
		admins.POST("/auth/refresh", h.adminRefresh)

		authenticated := admins.Group("/", h.verifyAdmin)
		{
			products := authenticated.Group("/products")
			{
				products.GET("/", h.getAllProducts)
				products.GET("/:id", h.getProductById)
				products.POST("/", h.createProduct)
				products.PUT("/:id", h.updateProduct)
				products.DELETE("/:id", h.deleteProduct)
				products.GET("/:id/reviews", h.getReviewsByProduct)
				products.POST("/:id/reviews", h.createReviewForProduct)
			}

			reviews := authenticated.Group("/reviews")
			{
				reviews.GET("/", h.getAllReviews)
				reviews.GET("/:id", h.getReviewById)
				reviews.POST("/", h.createReview)
				reviews.DELETE("/:id", h.deleteReview)
			}

			users := authenticated.Group("/users")
			{
				users.GET("/", getAllUsers)
				users.GET("/:id", getOneUser)
				users.POST("/", createUser)
				users.PUT("/:id", updateUser)
				users.DELETE("/:id", deleteUser)
			}
		}
	}
}

func (h *Handler) adminSignIn(context *gin.Context) {
	var adminDTO dto.AdminDTO
	err := context.BindJSON(&adminDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	admin, err := h.services.Admins.FindByCredentials(context, adminDTO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorResponse(context, http.StatusUnauthorized, "invalid admin credentials")
		} else {
			errorResponse(context, http.StatusInternalServerError, err.Error())
		}
		return
	}

	adminClaims := jwt.MapClaims{"adminID": admin.ID}
	authDetails, err := h.tokenProvider.CreateJWTSession(auth.CreateSessionInput{
		Fingerprint: adminDTO.Fingerprint,
		Claims:      adminClaims,
	})
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}
	successResponse(context, authDetails)
}

func (h *Handler) adminRefresh(context *gin.Context) {
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

func (h *Handler) verifyAdmin(context *gin.Context) {
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

	adminID, ok := tokenClaims["adminID"]
	if !ok {
		errorResponse(context, http.StatusForbidden, "admin endpoint is forbidden")
		return
	}

	context.Set("adminID", adminID)
}
