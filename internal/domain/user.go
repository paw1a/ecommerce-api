package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Image    string             `json:"image" bson:"image"`
	Password string             `json:"password" bson:"password"`
	Cart     Cart               `json:"cart" bson:"cart"`
	Orders   []Order            `json:"orders" bson:"orders"`
}

type Cart struct {
	TotalPrice float64    `json:"totalPrice" bson:"totalPrice"`
	CartItems  []CartItem `json:"cartItems" bson:"cartItems"`
}

type CartItem struct {
	ProductID primitive.ObjectID `json:"id" bson:"_id,omitempty"` //TODO: rename id
	Number    int                `json:"number" bson:"number"`
	ExpiresAt time.Time          `json:"expiresAt" bson:"expiresAt"`
}
