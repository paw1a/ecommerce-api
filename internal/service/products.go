package service

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"github.com/paw1a/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService struct {
	repo           repository.Products
	reviewsService Reviews
}

func (p *ProductsService) FindAll(ctx context.Context) ([]domain.Product, error) {
	products, err := p.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		productReviews, err := p.reviewsService.FindByProductID(ctx, product.ID)
		if err != nil {
			return nil, err
		}
		var ratingSum = 0
		for _, review := range productReviews {
			ratingSum += int(review.Rating)
		}
		product.TotalRating = float32(ratingSum / len(productReviews))
	}

	return products, nil
}

func (p *ProductsService) FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error) {
	product, err := p.repo.FindByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}
	productReviews, err := p.reviewsService.FindByProductID(ctx, product.ID)
	if err != nil {
		return domain.Product{}, err
	}
	var ratingSum = 0
	for _, review := range productReviews {
		ratingSum += int(review.Rating)
	}
	product.TotalRating = float32(ratingSum / len(productReviews))
	return product, nil
}

func (p *ProductsService) Create(ctx context.Context, product dto.CreateProductDTO) (domain.Product, error) {
	return p.repo.Create(ctx, domain.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Categories:  product.Categories,
	})
}

func (p *ProductsService) Update(ctx context.Context, productDTO dto.UpdateProductDTO, productID primitive.ObjectID) (domain.Product, error) {
	return p.repo.Update(ctx, dto.UpdateProductInput{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Categories:  productDTO.Categories,
	}, productID)
}

func (p *ProductsService) Delete(ctx context.Context, productID primitive.ObjectID) error {
	err := p.repo.Delete(ctx, productID)
	if err != nil {
		return err
	}
	return p.reviewsService.DeleteByProductID(ctx, productID)
}

func NewProductsService(repo repository.Products, reviewsService Reviews) *ProductsService {
	return &ProductsService{
		repo:           repo,
		reviewsService: reviewsService,
	}
}
