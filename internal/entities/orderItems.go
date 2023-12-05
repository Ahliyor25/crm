package entities

// OrderItems ...
type OrderItems struct {
	BaseGorm
	OrderID              uint64 `json:"order_id"`
	ProductID            uint64 `json:"product_id"`
	Quantity             uint64 `json:"quantity"`
	ProductBasePrice     uint64 `json:"product_base_price"`
	ProductCostPrice     uint64 `json:"product_cost_price"`
	ProductDiscountPrice uint64 `json:"product_discount_price"`
	// вспомагательная поле для orm
	Order   Order   `json:"-"`
	Product Product `json:"-"`
}
