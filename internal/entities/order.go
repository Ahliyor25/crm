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
	TrafficSourceID uint `json:"traffic_source_id"`
	// OrderItems
	OrderItems      []OrderItem `json:"order_items"`
	ShippingAddress string      `json:"shipping_address"`
	Discount        uint        `json:"discount"`
	SubTotal        float32     `json:"sub_total"`
	Total           float32     `json:"total"`

	// вспомагательная поле для orm
	Client        Client        `gorm:"foreignKey:ClientID"  json:"-"`
	User          User          `json:"-" gorm:"foreignKey:UserID" json:"-"`
	Status        Status        `gorm:"foreignKey:StatusID" json:"-"`
	TrafficSource TrafficSource `gorm:"foreignKey:TrafficSourceID" json:"-"`
}
