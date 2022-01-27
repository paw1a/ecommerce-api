package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewsService struct {
	repo repository.Reviews
}

func (r *ReviewsService) FindAll(ctx context.Context) ([]domain.Review, error) {
	return r.repo.FindAll(ctx)
}

func (r *ReviewsService) FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error) {
	return r.repo.FindByID(ctx, reviewID)
}

func (r *ReviewsService) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Review, error) {
	return r.repo.FindByUserID(ctx, userID)
}

func (r *ReviewsService) FindByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Review, error) {
	return r.repo.FindByProductID(ctx, productID)
}

func (r *ReviewsService) Create(ctx context.Context, review dto.CreateReviewInput) (domain.Review, error) {
	return r.repo.Create(ctx, domain.Review{
		UserID:    review.UserID,
		ProductID: review.ProductID,
		Text:      review.Text,
		Rating:    review.Rating,
	})
}

func (r *ReviewsService) Delete(ctx context.Context, reviewID primitive.ObjectID) error {
	return r.repo.Delete(ctx, reviewID)
}

func NewReviewsService(repo repository.Reviews) *ReviewsService {
	return &ReviewsService{
		repo: repo,
	}
}
