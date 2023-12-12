package entities

// Product ...
type Product struct {
	BaseGorm
	Sku           string  `json:"sku"`
	CostPrice     float32 `json:"cost_price"`
	BasePrice     float32 `json:"base_price"`
	Quantity      uint    `json:"quantity"`
	Discount      float32 `json:"discount"`
	MinQuantity   uint    `json:"min_quantity"`
	ProductInfoID uint    `json:"product_info_id"`
	
	// все что ниже - это связи с другими  таблицами
	ProductInfo ProductInfo `json:"-" gorm:"foreignKey:ProductInfoID"`
}
