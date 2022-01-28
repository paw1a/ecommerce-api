package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
	TotalRating float32            `json:"totalRating,omitempty" bson:"-"`
	Categories  []Category         `json:"categories" bson:"categories"` //TODO: make category independent entity, change type to Array(ObjectID)
}

type Category struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}
