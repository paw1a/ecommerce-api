package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}
