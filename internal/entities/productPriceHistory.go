package entities

// ProductPriceHistory ...
type ProductPriceHistory struct {
	BaseGorm
	// Идентификатор товара с которым связана цена
	ProductID uint    `json:"product_id"`
	CostPrice float32 `json:"cost_price"`
	BasePrice float32 `json:"base_price"`

	// Идентификатор пользователя, который изменил цену
	UserID uint `json:"user_id"`

	// вспомагательная поле для orm
	Product Product `json:"-"`
	User    User    `json:"-"`
}
