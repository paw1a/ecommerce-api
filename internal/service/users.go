package service

import "github.com/paw1a/http-server/internal/repository"

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}
