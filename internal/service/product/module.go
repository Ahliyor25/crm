package product

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/internal/storage/product"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBProduct)

// BProduct ...
type BProduct interface {
	Create(product entities.Product) (err error)
	Get(productID uint) (data entities.Product, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Product product.SProduct
	Logger  *logrus.Logger
}

// provider ...
type provider struct {
	product product.SProduct
	logger  *logrus.Logger
}

// NewBProduct ...
func NewBProduct(params Dependencies) BProduct {
	return &provider{
		product: params.Product,
		logger:  params.Logger,
	}
}
