package productpricehistory

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSProductPriceHistory)

// SProductPriceHistory ...
type SProductPriceHistory interface {
	Create(productPriceHistory entities.ProductPriceHistory) (err error)
	Get(target entities.ProductPriceHistory) (data entities.ProductPriceHistory, err error)
	Update(data entities.ProductPriceHistory) (err error)
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

// NewSProductPriceHistory ...
func NewSProductPriceHistory(params Dependencies) SProductPriceHistory {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
