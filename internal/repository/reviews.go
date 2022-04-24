package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ReviewsRepo struct {
	db *mongo.Collection
}

func (r ReviewsRepo) FindAll(ctx context.Context) ([]domain.Review, error) {
	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var reviewsArray []domain.Review
	err = cursor.All(ctx, &reviewsArray)
	return reviewsArray, err
}

func (r ReviewsRepo) FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error) {
	result := r.db.FindOne(ctx, bson.M{"_id": reviewID})

	var review domain.Review
	err := result.Decode(&review)

	return review, err
}

func (r ReviewsRepo) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Review, error) {
	cursor, err := r.db.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}

	var reviewsArray []domain.Review
	err = cursor.All(ctx, &reviewsArray)
	return reviewsArray, err
}

func (r ReviewsRepo) FindByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Review, error) {
	cursor, err := r.db.Find(ctx, bson.M{"productID": productID})
	if err != nil {
		return nil, err
	}

	var reviewsArray []domain.Review
	err = cursor.All(ctx, &reviewsArray)
	return reviewsArray, err
}

func (r ReviewsRepo) Create(ctx context.Context, review domain.Review) (domain.Review, error) {
	review.ID = primitive.NewObjectID()
	review.Date = time.Now()
	_, err := r.db.InsertOne(ctx, review)
	return review, err
}

func (r ReviewsRepo) Delete(ctx context.Context, reviewID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": reviewID})
	return err
}

func (r ReviewsRepo) DeleteByProductID(ctx context.Context, productID primitive.ObjectID) error {
	_, err := r.db.DeleteMany(ctx, bson.M{"productID": productID})
	return err
}

func NewReviewsRepo(db *mongo.Database) *ReviewsRepo {
	return &ReviewsRepo{
		db: db.Collection(reviewsCollection),
	}
}
