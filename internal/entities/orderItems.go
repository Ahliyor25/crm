package entities

// OrderItem ...
type OrderItem struct {
	BaseGorm
	OrderID              uint    `json:"order_id"`
	ProductID            uint    `json:"product_id"`
	Quantity             uint    `json:"quantity"`
	ProductBasePrice     float32 `json:"product_base_price"`
	ProductCostPrice     float32 `json:"product_cost_price"`
	ProductDiscountPrice float32 `json:"product_discount_price"`

	// вспомагательная поле для orm
	Order   Order   `gorm:"foreignkey:OrderID" json:"-"`
	Product Product `gorm:"foreignkey:ProductID" json:"-"`
}
