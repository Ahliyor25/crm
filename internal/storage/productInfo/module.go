package productinfo

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide()

// SUser ...
type SUser interface {
	Create(productInfo entities.ProductInfo) (err error)
	Get(target entities.ProductInfo) (data entities.ProductInfo, err error)
	GetMany(target entities.ProductInfo, currentProductInfoID uint) (productInfo []entities.ProductInfo, err error)
	Update(data entities.ProductInfo) (err error)
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

// NewSLogin ...
func NewSLogin(params Dependencies) SUser {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
