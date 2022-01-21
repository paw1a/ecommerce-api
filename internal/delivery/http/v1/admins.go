package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-in", h.adminSignIn)

		authenticated := admins.Group("/")
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
		newResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}

	admin, err := h.services.Admins.FindByCredentials(context, adminDTO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			newResponse(context, http.StatusUnauthorized, "Invalid admin credentials")
		} else {
			newResponse(context, http.StatusInternalServerError, err.Error())
		}
		return
	}

	token, err := h.tokenProvider.CreateJWTSession(admin, "")
	if err != nil {
		newResponse(context, http.StatusUnauthorized, err.Error())
		return
	}
	context.JSON(http.StatusOK, dataResponse{Data: token})
}
