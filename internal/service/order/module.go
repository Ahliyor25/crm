package order

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/internal/storage/order"
	"github.com/ahliyor25/crm/internal/storage/product"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBOrder)

// BOrder ...
type BOrder interface {
	Create(order entities.Order) (createdOrder entities.Order, err error)
	Get(orderID uint) (data entities.Order, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Order   order.SOrder
	Logger  *logrus.Logger
	Product product.SProduct
}

// provider ...
type provider struct {
	order   order.SOrder
	logger  *logrus.Logger
	product product.SProduct
}

// NewBOrder ...
func NewBOrder(params Dependencies) BOrder {
	return &provider{
		order:   params.Order,
		logger:  params.Logger,
		product: params.Product,
	}
}
