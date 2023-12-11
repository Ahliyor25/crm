package product

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSProduct)

// SUser ...
type SProduct interface {
	Create(product entities.Product) (err error)
	Get(target entities.Product) (data entities.Product, err error)
	GetMany(target entities.Product, currentProductID uint) (product []entities.Product, err error)
	Update(data entities.Product) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

type provider struct {
	logger   *logrus.Logger
	postgres *gorm.DB
}

// NewSProduct ...
func NewSProduct(params Dependencies) SProduct {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
