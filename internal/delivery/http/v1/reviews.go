package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// GetReviews godoc
// @Summary   Get all reviews
// @Tags      admin-reviews
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401     {object}  failure
// @Failure   404     {object}  failure
// @Failure   500     {object}  failure
// @Security  AdminAuth
// @Router    /admins/reviews [get]
func (h *Handler) getAllReviewsAdmin(context *gin.Context) {
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

// GetReviewById godoc
// @Summary   Get review by id
// @Tags      admin-reviews
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "review id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/reviews/{id} [get]
func (h *Handler) getReviewByIdAdmin(context *gin.Context) {
	id, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	review, err := h.services.Reviews.FindByID(context.Request.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorResponse(context, http.StatusInternalServerError,
				fmt.Sprintf("no reviews with id: %s", id.Hex()))
		} else {
			errorResponse(context, http.StatusInternalServerError, err.Error())
		}
		return
	}

	successResponse(context, review)
}

// CreateReview godoc
// @Summary   Create review
// @Tags      admin-reviews
// @Accept    json
// @Produce   json
// @Param     review  body      dto.CreateReviewDTOAdmin  true  "review"
// @Success   201     {object}  success
// @Failure   400     {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/reviews [post]
func (h *Handler) createReviewAdmin(context *gin.Context) {
	var reviewDTO dto.CreateReviewDTOAdmin
	err := context.BindJSON(&reviewDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}
	review, err := h.services.Reviews.Create(context.Request.Context(), dto.CreateReviewInput{
		UserID:    reviewDTO.UserID,
		ProductID: reviewDTO.ProductID,
		Text:      reviewDTO.Text,
		Rating:    reviewDTO.Rating,
	})

	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, review)
}

// DeleteReview godoc
// @Summary   Delete review
// @Tags      admin-reviews
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "review id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/reviews/{id} [delete]
func (h *Handler) deleteReviewAdmin(context *gin.Context) {
	reviewID, err := getIdFromPath(context, "id")
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
