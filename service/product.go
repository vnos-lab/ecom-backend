package service

import (
	"context"
	"ecom/domain"
	"ecom/models"

	"go.uber.org/zap"
)

type productService struct {
	productRepo domain.ProductRepository
	logger      *zap.Logger
}

func NewProductService(productRepo domain.ProductRepository, logger *zap.Logger) domain.ProductService {
	return &productService{
		productRepo: productRepo,
		logger:      logger,
	}
}

func (p *productService) Create(ctx context.Context, req domain.CreateProductRequest) (*models.Product, error) {
	return p.productRepo.Create(ctx, models.Product{
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
		Thumbnail:   req.Thumbnail,
		Images:      req.Images,
		IsNew:       req.IsNew,
	})
}

func (p *productService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return p.productRepo.GetByID(ctx, id)
}
