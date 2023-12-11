package productinfo

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSProductInfo)

// SProductInfo ...
type SProductInfo interface {
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

// NewSProductInfo ...
func NewSProductInfo(params Dependencies) SProductInfo {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
