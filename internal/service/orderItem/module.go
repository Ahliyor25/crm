package orderItem

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/internal/storage/orderItem"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBOrderItem)

// BOrderItem ...
type BOrderItem interface {
	Create(orderItem entities.OrderItem) (err error)
	Get(orderItemID uint) (data entities.OrderItem, err error)
	GetList(orderId uint) (data []entities.OrderItem, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	OrderItem orderItem.SOrderItem
	Logger    *logrus.Logger
}

// provider ...
type provider struct {
	orderItem orderItem.SOrderItem
	logger    *logrus.Logger
}

// NewBOrderItem ...
func NewBOrderItem(params Dependencies) BOrderItem {
	return &provider{
		orderItem: params.OrderItem,
		logger:    params.Logger,
	}
}
