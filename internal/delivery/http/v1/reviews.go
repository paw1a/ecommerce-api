package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) getAllReviews(context *gin.Context) {
	reviews, err := h.services.Reviews.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	reviewsArray := make([]domain.Review, len(reviews))
	if reviews != nil {
		reviewsArray = reviews
	}

	successResponse(context, reviewsArray)
}

func (h *Handler) getReviewById(context *gin.Context) {
	id, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	review, err := h.services.Reviews.FindByID(context.Request.Context(), id)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, review)
}

func (h *Handler) createReview(context *gin.Context) {
	var reviewDTO dto.CreateReviewDTO
	err := context.BindJSON(&reviewDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}
	review, err := h.services.Reviews.Create(context.Request.Context(), dto.CreateReviewInput{
		UserID:    primitive.ObjectID{},
		ProductID: primitive.ObjectID{},
		Text:      reviewDTO.Text,
		Rating:    reviewDTO.Rating,
	})

	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, review)
}

func (h *Handler) deleteReview(context *gin.Context) {
	reviewID, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Reviews.Delete(context, reviewID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func (h *Handler) getReviewsByProduct(context *gin.Context) {
	productID, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	reviews, err := h.services.Reviews.FindByProductID(context, productID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, reviews)
}

func (h *Handler) createReviewForProduct(context *gin.Context) {
	productID, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var reviewDTO dto.CreateReviewDTO
	err = context.BindJSON(&reviewDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}

	review, err := h.services.Reviews.Create(context, dto.CreateReviewInput{
		UserID:    primitive.ObjectID{},
		ProductID: productID,
		Text:      reviewDTO.Text,
		Rating:    reviewDTO.Rating,
	})

	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, review)
}
