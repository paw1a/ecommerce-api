package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateReviewDTO struct {
	Text   string `json:"text"`
	Rating int8   `json:"rating"`
}

type CreateReviewInput struct {
	UserID    primitive.ObjectID `json:"userID"`
	ProductID primitive.ObjectID `json:"productID"`
	Text      string             `json:"text"`
	Rating    int8               `json:"rating"`
}
