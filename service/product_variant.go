package service

import (
	"context"
	"ecom/domain"
	"ecom/models"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type productVariantService struct {
	variantRepo domain.ProductVariantRepository
	logger      *zap.Logger
}

func NewVariantService(variantRepo domain.ProductVariantRepository, logger *zap.Logger) domain.ProductVariantService {
	return &productVariantService{variantRepo, logger}
}

func (v *productVariantService) Create(context context.Context, req domain.CreateProductVariantRequest) (*models.ProductVariant, error) {
	r, err := v.variantRepo.Create(context, models.ProductVariant{
		Price:     req.Price,
		Stock:     req.Stock,
		SKU:       req.SKU,
		Variant:   req.Variant,
		ProductID: req.ProductID,
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (v *productVariantService) Delete(context context.Context, id string) (err error) {
	return v.variantRepo.Delete(context, id)
}

func (v *productVariantService) Update(context context.Context, req domain.UpdateProductVariantRequest, id uuid.UUID) (*models.ProductVariant, error) {
	r, err := v.variantRepo.Update(context, id, models.ProductVariant{
		Price:     req.Price,
		Stock:     req.Stock,
		SKU:       req.SKU,
		Variant:   req.Variant,
		ProductID: req.ProductID,
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}
