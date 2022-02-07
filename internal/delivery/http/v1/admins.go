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
				products.GET("/", h.getAllProductsAdmin)
				products.GET("/:id", h.getProductByIdAdmin)
				products.POST("/", h.createProductAdmin)
				products.PUT("/:id", h.updateProductAdmin)
				products.DELETE("/:id", h.deleteProductAdmin)
				products.GET("/:id/reviews", h.getProductReviewsAdmin)
			}

			reviews := authenticated.Group("/reviews")
			{
				reviews.GET("/", h.getAllReviewsAdmin)
				reviews.GET("/:id", h.getReviewByIdAdmin)
				reviews.POST("/", h.createReviewAdmin)
				reviews.DELETE("/:id", h.deleteReviewAdmin)
			}

			users := authenticated.Group("/users")
			{
				users.GET("/", h.getAllUsersAdmin)
				users.GET("/:id", h.getUserByIdAdmin)
				users.POST("/", h.createUserAdmin)
				users.PUT("/:id", h.updateUserAdmin)
				users.DELETE("/:id", h.deleteUserAdmin)
			}

			cart := authenticated.Group("/carts")
			{
				cart.GET("/", h.getAllCartsAdmin)
				cart.GET("/:id", h.getCartByIdAdmin)
				cart.DELETE("/:id", h.deleteCartAdmin)
			}

			orders := authenticated.Group("/orders")
			{
				orders.GET("/", h.getAllOrdersAdmin)
				orders.PUT("/:id", h.updateOrderAdmin)
			}
		}
	}
}

// AdminSignIn godoc
// @Summary  Admin sign-in
// @Tags     admin-auth
// @Accept   json
// @Produce  json
// @Param    admin  body      dto.SignInDTO  true  "admin credentials"
// @Success  200    {object}  auth.AuthDetails
// @Failure  400    {object}  failure
// @Failure  401    {object}  failure
// @Failure  404    {object}  failure
// @Failure  500    {object}  failure
// @Router   /admins/auth/sign-in [post]
func (h *Handler) adminSignIn(context *gin.Context) {
	var signInDTO dto.SignInDTO
	err := context.BindJSON(&signInDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	admin, err := h.services.Admins.FindByCredentials(context, signInDTO)
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
		Fingerprint: signInDTO.Fingerprint,
		Claims:      adminClaims,
	})
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}
	successResponse(context, authDetails)
}

// AdminRefresh godoc
// @Summary  Admin refresh token
// @Tags     admin-auth
// @Accept   json
// @Produce  json
// @Param    refreshInput  body      auth.RefreshInput  true  "refresh data"
// @Success  200           {object}  auth.AuthDetails
// @Failure  400           {object}  failure
// @Failure  401           {object}  failure
// @Failure  404           {object}  failure
// @Failure  500           {object}  failure
// @Router   /admins/auth/refresh [post]
func (h *Handler) adminRefresh(context *gin.Context) {
	h.refreshToken(context)
}

func (h *Handler) verifyAdmin(context *gin.Context) {
	h.verifyToken(context, "adminID")
}
