package postgres

import (
	"context"
	"ecom/api_errors"
	"ecom/domain"
	"ecom/infrastructure/db"
	"ecom/models"
	"ecom/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type productRepository struct {
	*db.Database
	logger *zap.Logger
}

func NewProductRepository(db *db.Database, logger *zap.Logger) domain.ProductRepository {
	utils.MustHaveDb(db)
	return &productRepository{db, logger}
}

func (p *productRepository) Create(context context.Context, product models.Product) (*models.Product, error) {
	args := []interface{}{
		product.Name,
		product.Description,
		product.Slug,
		product.Thumbnail,
		product.Images,
		product.IsNew,
		utils.GetUserIDFromContext(context),
	}

	sql, _, _ := utils.Psql().Insert(models.Product{}.TableName()).Columns(
		"name",
		"description",
		"slug",
		"thumbnail",
		"images",
		"is_new",
		"updater_id",
		// "category_id",
		// "campaign_id",
	).Values(args...).Suffix("RETURNING id, created_at, updated_at").ToSql()

	err := p.DB.QueryRow(sql, args...).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &product, nil
}

func (p *productRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product

	r := []domain.GetByIDQueryResult{}

	sql, _, _ := utils.Psql().Select(`products.*,product_variants.id "variant.id",product_variants.updater_id "variant.updater_id",
	product_variants.created_at "variant.created_at",product_variants.updated_at "variant.updated_at",
	product_variants.price "variant.price",product_variants.stock "variant.stock",product_variants.sku "variant.sku",
	product_variants.variant "variant.variant",product_variants.product_id "variant.product_id"`).
		Distinct().
		From(models.Product{}.TableName()).
		LeftJoin("product_variants ON product_variants.product_id = products.id").
		Where("products.id = ?").
		ToSql()
	err := p.DB.SelectContext(ctx, &r, sql, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(r) == 0 {
		return nil, errors.WithStack(api_errors.ErrProductNotFound)
	}

	getByIDQueryResultToProduct(r, &product)

	return &product, nil
}

// đặt tên cho tôi function này lại cho đúng với mục đích của nó

func getByIDQueryResultToProduct(r []domain.GetByIDQueryResult, product *models.Product) {
	for _, v := range r {
		product.ID = v.ID
		product.Name = v.Name
		product.Description = v.Description
		product.Slug = v.Slug
		product.Images = v.Images
		product.Thumbnail = v.Thumbnail
		product.IsNew = v.IsNew
		product.CategoryID = v.CategoryID
		product.CampaignID = v.CampaignID
		product.TotalSold = v.TotalSold
		product.Rating = v.Rating
		product.CreatedAt = v.CreatedAt
		product.UpdatedAt = v.UpdatedAt
		product.Variants = append(product.Variants, &models.ProductVariant{
			BaseModel: models.BaseModel{
				ID:        v.Variant.ID,
				UpdaterID: v.Variant.UpdaterID,
				CreatedAt: v.Variant.CreatedAt,
				UpdatedAt: v.Variant.UpdatedAt,
			},
			Price:     v.Variant.Price,
			Stock:     v.Variant.Stock,
			SKU:       v.Variant.SKU,
			Variant:   v.Variant.Variant,
			ProductID: v.Variant.ProductID,
		})
	}
}
