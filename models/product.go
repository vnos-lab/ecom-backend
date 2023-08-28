package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

type Product struct {
	BaseModel
	Name        string            `json:"name" db:"name"`
	Description string            `json:"description" db:"description"`
	Slug        string            `json:"slug" db:"slug"`
	Images      Images            `json:"images" db:"images"`
	Thumbnail   string            `json:"thumbnail" db:"thumbnail"`
	IsNew       bool              `json:"is_new" db:"is_new"`
	CategoryID  *int              `json:"category_id" db:"category_id"`
	CampaignID  *int              `json:"campaign_id" db:"campaign_id"`
	TotalSold   int               `json:"total_sold" db:"total_sold"`
	Rating      float64           `json:"rating" db:"rating"`
	Variants    []*ProductVariant `json:"variants" db:"variants"`
}

type Images []string

func (s *Images) Scan(src interface{}) (err error) {
	var images []string

	err = json.Unmarshal(src.([]byte), &images)
	if err != nil {
		return errors.New("failed to unmarshal images")
	}

	*s = images
	return nil
}

func (s Images) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (p Product) TableName() string {
	return "products"
}
