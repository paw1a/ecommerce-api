package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Name string `json:"name"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
}
