package orderItem

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSOrderItem)

// SOrderItem ...
type SOrderItem interface {
	Create(orderItem entities.OrderItem) (err error)
	Get(target entities.OrderItem) (data entities.OrderItem, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

// provider ...
type provider struct {
	logger   *logrus.Logger
	postgres *gorm.DB
}

// NewSOrderItem ...
func NewSOrderItem(params Dependencies) SOrderItem {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
