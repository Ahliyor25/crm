package entities

// OrderItem ...
type OrderItem struct {
	BaseGorm
	OrderID              uint    `json:"order_id"`
	ProductID            uint    `json:"product_id"`
	Quantity             uint    `json:"quantity"`
	ProductName          string  `json:"product_name"`
	ProductBasePrice     float32 `json:"product_base_price"`
	ProductCostPrice     float32 `json:"product_cost_price"`
	ProductDiscountPrice float32 `json:"product_discount_price" gorm:"default:0" validate:"min=0,max=100"`

	// вспомагательная поле для orm
	Order   Order   `gorm:"foreignKey:OrderID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"-"`
}
