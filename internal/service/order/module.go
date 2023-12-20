package order

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/internal/storage/order"
	"github.com/ahliyor25/crm/internal/storage/product"
	productinfo "github.com/ahliyor25/crm/internal/storage/productInfo"
	"github.com/ahliyor25/crm/internal/storage/status"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBOrder)

// BOrder ...
type BOrder interface {
	Create(order entities.Order) (createdOrder entities.Order, err error)
	Get(orderID uint) (data entities.Order, err error)
	Update(order entities.Order) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Order       order.SOrder
	Logger      *logrus.Logger
	Product     product.SProduct
	ProductInfo productinfo.SProductInfo
	Status      status.SStatus
}

// provider ...
type provider struct {
	order       order.SOrder
	logger      *logrus.Logger
	product     product.SProduct
	status      status.SStatus
	productInfo productinfo.SProductInfo
}

// NewBOrder ...
func NewBOrder(params Dependencies) BOrder {
	return &provider{
		order:       params.Order,
		logger:      params.Logger,
		product:     params.Product,
		status:      params.Status,
		productInfo: params.ProductInfo,
	}
}
