package order

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSOrder)

// SOrder ...
type SOrder interface {
	Create(order entities.Order) (err error)
	Get(target entities.Order) (data entities.Order, err error)
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

// NewSOrder ...
func NewSOrder(params Dependencies) SOrder {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
