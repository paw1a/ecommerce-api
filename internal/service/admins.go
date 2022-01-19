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

func (a *AdminsService) FindByCredentials(ctx context.Context, adminDTO dto.AdminDTO) (domain.Admin, error) {
	return a.repo.FindByCredentials(ctx, adminDTO.Email, adminDTO.Password)
}

func NewAdminsService(repo repository.Admins) *AdminsService {
	return &AdminsService{
		repo: repo,
	}
}
