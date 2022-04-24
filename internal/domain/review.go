package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Review struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID"`
	ProductID primitive.ObjectID `json:"productID" bson:"productID"`
	Text      string             `json:"text" bson:"text"`
	Rating    int8               `json:"rating" bson:"rating"`
	Username  string             `json:"username" bson:"username"`
	Date      time.Time          `json:"date" bson:"date"`
}
