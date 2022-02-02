package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Name   string              `json:"name"`
	CartID *primitive.ObjectID `json:"cartID"`
}

type UpdateUserInput struct {
	Name   string              `json:"name"`
	CartID *primitive.ObjectID `json:"cartID"`
}

type SignUpDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
