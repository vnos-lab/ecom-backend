package domain

import (
	"context"
	"ecom/models"

	uuid "github.com/satori/go.uuid"
)

type ProductVariantRepository interface {
	GetByProductID(ctx context.Context, productID string) (res []*models.ProductVariant, err error)
	Create(ctx context.Context, variant models.ProductVariant) (res *models.ProductVariant, err error)
	Update(ctx context.Context, id uuid.UUID, variant models.ProductVariant) (res *models.ProductVariant, err error)
	Delete(ctx context.Context, id string) (err error)
}

type ProductVariantService interface {
	Create(ctx context.Context, req CreateProductVariantRequest) (res *models.ProductVariant, err error)
	Update(ctx context.Context, req UpdateProductVariantRequest, id uuid.UUID) (res *models.ProductVariant, err error)
	Delete(ctx context.Context, id string) (err error)
}

type CreateProductVariantRequest struct {
	Price     int               `json:"price" db:"price"`
	Stock     int               `json:"stock" db:"stock"`
	SKU       string            `json:"sku" db:"sku"`
	Variant   map[string]string `json:"variant" db:"variant"`
	ProductID string            `json:"product_id" db:"product_id"`
}

type UpdateProductVariantRequest struct {
	CreateProductVariantRequest
}
