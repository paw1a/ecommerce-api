package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService struct {
	repo repository.Products
}

func (p ProductsService) FindAll(ctx context.Context) ([]domain.Product, error) {
	return p.repo.FindAll(ctx)
}

func (p ProductsService) FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error) {
	return p.repo.FindByID(ctx, productID)
}

func (p ProductsService) Create(ctx context.Context, product dto.CreateProductDTO) (domain.Product, error) {
	return p.repo.Create(ctx, domain.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Categories:  product.Categories,
	})
}

func (p ProductsService) Update(ctx context.Context, productDTO dto.UpdateProductDTO, productID primitive.ObjectID) (domain.Product, error) {
	return p.repo.Update(ctx, dto.UpdateProductInput{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		TotalRating: productDTO.TotalRating,
		Categories:  productDTO.Categories,
	}, productID)
}

func (p ProductsService) Delete(ctx context.Context, productID primitive.ObjectID) error {
	return p.repo.Delete(ctx, productID)
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{
		repo: repo,
	}
}
