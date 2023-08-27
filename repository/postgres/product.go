package postgres

import (
	"context"
	"ecom/domain"
	"ecom/infrastructure/db"
	"ecom/models"
	"ecom/utils"

	"go.uber.org/zap"
)

type productRepository struct {
	*db.Database
	logger *zap.Logger
}

func NewProductRepository(db *db.Database, logger *zap.Logger) domain.ProductRepository {
	utils.ErrNilDb(db)
	return &productRepository{db, logger}
}

func (p *productRepository) Create(context context.Context, product models.Product) (*models.Product, error) {
	args := []interface{}{
		product.Name,
		product.Description,
		product.Price,
		product.Slug,
		product.Thumbnail,
		product.Images,
		product.IsNew,
		// product.CategoryID,
		// product.CampaignID,
		product.TotalSold,
		product.SKU,
	}

	sql, _, _ := utils.Psql().Insert(models.Product{}.TableName()).Columns(
		"name",
		"description",
		"price",
		"slug",
		"thumbnail",
		"images",
		"is_new",
		// "category_id",
		// "campaign_id",
		"total_sold",
		"sku",
	).Values(args...).Suffix("RETURNING id, created_at, updated_at").ToSql()

	err := p.DB.QueryRow(sql, args...).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		p.logger.Error("Error when insert product", zap.Error(err))
		return nil, err
	}
	return &product, nil
}
