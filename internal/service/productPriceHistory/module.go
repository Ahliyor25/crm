package productpricehistory

import (
	"github.com/ahliyor25/crm/internal/entities"
	productpricehistory "github.com/ahliyor25/crm/internal/storage/productPriceHistory"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBProductPriceHistory)

// BProductPriceHistory ...
type BProductPriceHistory interface {
	Create(productPriceHistory entities.ProductPriceHistory) (err error)
	Get(productPriceHistoryID uint) (data entities.ProductPriceHistory, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	ProductPriceHistory productpricehistory.SProductPriceHistory
	Logger              *logrus.Logger
}

// provider ...
type provider struct {
	productPriceHistory productpricehistory.SProductPriceHistory
	logger              *logrus.Logger
}

// NewBProductPriceHistory ...
func NewBProductPriceHistory(params Dependencies) BProductPriceHistory {
	return &provider{
		productPriceHistory: params.ProductPriceHistory,
		logger:              params.Logger,
	}
}
