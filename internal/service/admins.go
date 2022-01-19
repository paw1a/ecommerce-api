package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/repository"
)

type AdminsService struct {
	repo repository.Admins
}

func (a *AdminsService) FindByCredentials(ctx context.Context, email string, password string) (domain.Admin, error) {
	return a.repo.FindByCredentials(ctx, email, password)
}

func NewAdminsService(repo repository.Admins) *AdminsService {
	return &AdminsService{
		repo: repo,
	}
}
