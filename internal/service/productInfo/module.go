package productinfo

import (
	"github.com/ahliyor25/crm/internal/entities"
	productinfo "github.com/ahliyor25/crm/internal/storage/productInfo"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBProductInfo)

// BProductInfo ...
type BProductInfo interface {
	Create(productInfo entities.ProductInfo) (err error)
	Get(productInfoID uint) (data entities.ProductInfo, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	ProductInfo productinfo.SProductInfo
	Logger      *logrus.Logger
}

// provider ...
type provider struct {
	productInfo productinfo.SProductInfo
	logger      *logrus.Logger
}

// NewBProductInfo ...
func NewBProductInfo(params Dependencies) BProductInfo {
	return &provider{
		productInfo: params.ProductInfo,
		logger:      params.Logger,
	}
}
