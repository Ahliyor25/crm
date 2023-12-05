package entities

// Order ...
type Order struct {
	BaseGorm
	ClientID uint `json:"client_id"`
	// Идентификатор пользователя, который создал заказ
	UserID uint `json:"user_id"`
	// Идентификатор статуса заказа
	StatusID uint `json:"status_id"`
	// Идентификатор трафика заказа
	TrafficSourceID uint   `json:"traffic_source_id"`
	ShippingAddress string `json:"shipping_address"`
	Discount        uint   `json:"discount"`
	SubTotal        uint   `json:"sub_total"`
	Total           uint   `json:"total"`

	// вспомагательная поле для orm
	Client        Client        `json:"-"`
	User          User          `json:"-"`
	Status        Status        `json:"-"`
	TrafficSource TrafficSource `json:"-"`
}
