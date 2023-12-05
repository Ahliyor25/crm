package entities

import "time"

// ProductPublication ...
type ProductPublication struct {
	ProductID uint    `json:"product_id"`
	SourceID  uint    `json:"source_id"`
	Link      string  `json:"link"`
	Price     float32 `json:"price"`
	Quantity  uint    `json:"quantity"`
	StatusID  uint    `json:"status_id"`

	// вспомагательная поле для orm
	Product Product       `json:"-"`
	Source  TrafficSource `json:"-"`
	Status  Status        `json:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
