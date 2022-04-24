package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"time"
)

type ReviewsService struct {
	repo        repository.Reviews
	redisClient *redis.Client
	userService Users
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

func (r *ReviewsService) Create(ctx context.Context, reviewDTO dto.CreateReviewInput) (domain.Review, error) {
	user, err := r.userService.FindByID(ctx, reviewDTO.UserID)
	if err != nil {
		return domain.Review{}, err
	}

	review, err := r.repo.Create(ctx, domain.Review{
		UserID:    reviewDTO.UserID,
		ProductID: reviewDTO.ProductID,
		Text:      reviewDTO.Text,
		Rating:    reviewDTO.Rating,
		Username:  user.Name,
	})
	if err != nil {
		return domain.Review{}, err
	}

	_, err = r.calculateProductRating(ctx, review.ProductID)

	return review, err
}

func (r *ReviewsService) calculateProductRating(ctx context.Context, productID primitive.ObjectID) (float64, error) {
	productReviews, err := r.FindByProductID(ctx, productID)
	if err != nil {
		return 0.0, err
	}

	var ratingSum = 0
	for _, review := range productReviews {
		ratingSum += int(review.Rating)
	}

	var rating float64
	if len(productReviews) != 0 {
		value := float64(ratingSum) / float64(len(productReviews))
		rating = math.Floor(value*10) / 10
	}

	err = r.redisClient.Set(productID.Hex(), rating, time.Hour*24*7).Err()
	return rating, err
}

func (r *ReviewsService) GetTotalReviewRating(ctx context.Context, productID primitive.ObjectID) (float64, error) {
	var rating float64
	cachedRating, err := r.redisClient.Get(productID.Hex()).Float64()
	if err != nil {
		rating, err = r.calculateProductRating(ctx, productID)
	} else {
		rating = cachedRating
	}

	return rating, err
}

func (r *ReviewsService) Delete(ctx context.Context, reviewID primitive.ObjectID) error {
	review, err := r.repo.FindByID(ctx, reviewID)
	if err != nil {
		return fmt.Errorf("no review with id: %s", reviewID.Hex())
	}
	err = r.repo.Delete(ctx, reviewID)
	if err != nil {
		return fmt.Errorf("failed to delete review with id: %s", reviewID)
	}

	_, err = r.calculateProductRating(ctx, review.ProductID)

	return err
}

func (r *ReviewsService) DeleteByProductID(ctx context.Context, productID primitive.ObjectID) error {
	err := r.repo.DeleteByProductID(ctx, productID)
	if err != nil {
		return fmt.Errorf("failed to delete review with productID: %s", productID)
	}

	_, err = r.calculateProductRating(ctx, productID)

	return err
}

func NewReviewsService(repo repository.Reviews, redisClient *redis.Client) *ReviewsService {
	return &ReviewsService{
		repo:        repo,
		redisClient: redisClient,
	}
}
