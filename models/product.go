package models

type Product struct {
	BaseModel
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Price       int      `json:"price" db:"price"`
	Slug        string   `json:"slug" db:"slug"`
	Thumbnail   string   `json:"thumbnail" db:"thumbnail"`
	Images      []string `json:"images" db:"images"`
	IsNew       bool     `json:"is_new" db:"is_new"`
	CategoryID  int      `json:"category_id" db:"category_id"`
	CampaignID  int      `json:"campaign_id" db:"campaign_id"`
	TotalSold   int      `json:"total_sold" db:"total_sold"`
	SKU         string   `json:"sku" db:"sku"`
}

func (p Product) TableName() string {
	return "products"
}
