package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cart struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	ExpireAt   time.Time          `json:"expireAt" bson:"expireAt"`
	TotalPrice float64            `json:"totalPrice" bson:"-"`
	CartItems  []CartItem         `json:"cartItems" bson:"cartItems"`
}

type CartItem struct {
	ProductID primitive.ObjectID `json:"productID" bson:"productID"`
	Quantity  int64              `json:"quantity" bson:"quantity"`
}
