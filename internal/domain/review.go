package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID"`
	ProductID primitive.ObjectID `json:"productID" bson:"productID"`
	Text      string             `json:"text" bson:"text"`
	Rating    int8               `json:"rating" bson:"rating"`
}
