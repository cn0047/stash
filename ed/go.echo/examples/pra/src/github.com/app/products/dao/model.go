package dao

// ProductModel - product model structure.
// For represent any product in the project this structure must be used.
type ProductModel struct {
	BrandName   string  `json:"brand_name"`
	DateCreated string  `json:"date_created"`
	Gender      string  `json:"gender"`
	ID          string  `json:"id"`
	InStock     bool    `json:"in_stock"`
	Name        string  `json:"name"`
	ProductURL  string  `json:"product_url"`
	Price       float32 `json:"price"`
	RetailPrice float32 `json:"retail_price"`
	StoreName   string  `json:"store_name"`
}
