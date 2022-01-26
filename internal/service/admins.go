package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
)

type AdminsService struct {
	repo repository.Admins
}

func (a *AdminsService) FindByCredentials(ctx context.Context, signInDTO dto.SignInDTO) (domain.Admin, error) {
	return a.repo.FindByCredentials(ctx, signInDTO.Email, signInDTO.Password)
}

func NewAdminsService(repo repository.Admins) *AdminsService {
	return &AdminsService{
		repo: repo,
	}
}
