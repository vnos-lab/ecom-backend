package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ProductVariant struct {
	BaseModel
	Price     int     `json:"price" db:"price"`
	Stock     int     `json:"stock" db:"stock"`
	SKU       string  `json:"sku" db:"sku"`
	Variant   Variant `json:"variant" db:"variant"`
	ProductID string  `json:"product_id" db:"product_id"`
}

type Variant map[string]string

func (s *Variant) Scan(src interface{}) (err error) {
	var variant map[string]string

	err = json.Unmarshal(src.([]byte), &variant)
	if err != nil {
		return errors.New("failed to unmarshal variant")
	}

	*s = variant
	return nil
}

func (s Variant) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (ProductVariant) TableName() string {
	return "product_variants"
}
