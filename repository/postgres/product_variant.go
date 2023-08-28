package postgres

import (
	"context"
	"ecom/domain"
	"ecom/infrastructure/db"
	"ecom/models"
	"ecom/utils"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type productVariantRepository struct {
	db     *db.Database
	logger *zap.Logger
}

func NewVariantRepository(db *db.Database, logger *zap.Logger) domain.ProductVariantRepository {
	return &productVariantRepository{db, logger}
}

func (v *productVariantRepository) Create(ctx context.Context, variant models.ProductVariant) (*models.ProductVariant, error) {
	e := new(models.ProductVariant)

	fmt.Println(variant)

	sql, _, _ := utils.Psql().Insert(models.ProductVariant{}.TableName()).Columns("product_id", "price", "stock", "sku", "variant", "updater_id").Values("?", "?", "?", "?", "?", "?").Suffix("RETURNING *").ToSql()
	err := v.db.QueryRowxContext(ctx, sql, variant.ProductID, variant.Price, variant.Stock, variant.SKU, variant.Variant, utils.GetUserIDFromContext(ctx)).StructScan(e)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return e, nil
}

func (v *productVariantRepository) Delete(ctx context.Context, id string) (err error) {
	sql, _, _ := utils.Psql().Delete(models.ProductVariant{}.TableName()).Where("id = ?").ToSql()

	_, err = v.db.ExecContext(ctx, sql, id)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (v *productVariantRepository) GetByProductID(ctx context.Context, productID string) (res []*models.ProductVariant, err error) {
	sql, _, _ := utils.Psql().Select("*").From(models.ProductVariant{}.TableName()).Where("product_id = ?").ToSql()

	err = v.db.SelectContext(ctx, &res, sql, productID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

func (v *productVariantRepository) Update(ctx context.Context, id uuid.UUID, variant models.ProductVariant) (*models.ProductVariant, error) {
	e := models.ProductVariant{}
	sql, _, _ := utils.Psql().Update(models.ProductVariant{}.TableName()).
		Set("price", "?").
		Set("stock", "?").
		Set("sku", "?").
		Set("variant", "?").
		Set("updater_id", "?").
		Where("id = ?").
		Suffix("RETURNING *").ToSql()
	err := v.db.QueryRowxContext(ctx, sql, variant.Price, variant.Stock, variant.SKU, variant.Variant, utils.GetUserIDFromContext(ctx), id).StructScan(&e)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &e, nil
}
