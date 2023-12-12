package entities

// ProductPublication ...
type ProductPublication struct {
	BaseGorm
	ProductID uint    `json:"product_id"`
	SourceID  uint    `json:"source_id"`
	Link      string  `json:"link"`
	Price     float32 `json:"price"`
	Quantity  uint    `json:"quantity"`
	StatusID  uint    `json:"status_id"`

	// вспомагательная поле для orm
	Product Product       `json:"-" gorm:"foreignKey:ProductID"`
	Source  TrafficSource `json:"-" gorm:"foreignKey:SourceID"`
	Status  Status        `json:"-" gorm:"foreignKey:StatusID"`
}
