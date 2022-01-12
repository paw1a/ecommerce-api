package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"net/http"
)

func (h *Handler) initReviewsRoutes(api *gin.RouterGroup) {
	users := api.Group("/reviews")
	{
		users.GET("/", h.getAllReviews)
		users.GET("/:id", h.getReviewById)
		users.POST("/", h.createReview)
		users.DELETE("/:id", h.deleteReview)
	}
}

func (h *Handler) getAllReviews(context *gin.Context) {
	reviews, err := h.services.Reviews.FindAll(context.Request.Context())
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	reviewsArray := make([]domain.Review, len(reviews))
	if reviews != nil {
		reviewsArray = reviews
	}

	context.JSON(http.StatusOK, dataResponse{Data: reviewsArray})
}

func (h *Handler) getReviewById(context *gin.Context) {
	id, err := parseIdFromPath(context, "id")
	if err != nil {
		newResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	review, err := h.services.Reviews.FindByID(context.Request.Context(), id)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, dataResponse{Data: review})
}

func (h *Handler) createReview(context *gin.Context) {
	var reviewDTO dto.CreateReviewDTO
	err := context.BindJSON(&reviewDTO)
	if err != nil {
		newResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}
	review, err := h.services.Reviews.Create(context.Request.Context(), reviewDTO)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, dataResponse{Data: review})
}

func (h *Handler) deleteReview(context *gin.Context) {
	reviewID, err := parseIdFromPath(context, "id")
	if err != nil {
		newResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Reviews.Delete(context, reviewID)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}
