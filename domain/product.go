package domain

import (
	"context"
	"ecom/models"
)

type ProductRepository interface {
	Create(ctx context.Context, product models.Product) (*models.Product, error)
}

type ProductService interface {
	Create(ctx context.Context, req CreateProductRequest) (*models.Product, error)
}

type CreateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Slug        string   `json:"slug" binding:"required"`
	Thumbnail   string   `json:"thumbnail" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	IsNew       bool     `json:"is_new" binding:"required"`
	// CategoryID int      `json:"category_id" binding:"required"`
	// CampaignID int      `json:"campaign_id" binding:"required"`
	// TotalSold int    `json:"total_sold" binding:"required"`
	SKU string `json:"sku" binding:"required"`
}
