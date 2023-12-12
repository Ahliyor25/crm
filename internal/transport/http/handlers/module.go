package handlers

import (
	"github.com/ahliyor25/crm/internal/service/auth"
	"github.com/ahliyor25/crm/internal/service/client"
	"github.com/ahliyor25/crm/internal/service/order"
	"github.com/ahliyor25/crm/internal/service/orderItem"
	"github.com/ahliyor25/crm/internal/service/product"
	productinfo "github.com/ahliyor25/crm/internal/service/productInfo"
	productpricehistory "github.com/ahliyor25/crm/internal/service/productPriceHistory"
	"github.com/ahliyor25/crm/internal/service/status"
	trafficsource "github.com/ahliyor25/crm/internal/service/trafficSource"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger              *logrus.Logger
	Auth                auth.BAuth
	Client              client.BClient
	ProductInfo         productinfo.BProductInfo
	Product             product.BProduct
	ProductPriceHistory productpricehistory.BProductPriceHistory
	Status              status.BStatus
	TrafficSource       trafficsource.BTrafficSource
	Order               order.BOrder
	OrderItem           orderItem.BOrderItem
}

// Handler ...
type Handler struct {
	logger              *logrus.Logger
	auth                auth.BAuth
	client              client.BClient
	productInfo         productinfo.BProductInfo
	product             product.BProduct
	productPriceHistory productpricehistory.BProductPriceHistory
	status              status.BStatus
	trafficSource       trafficsource.BTrafficSource
	order               order.BOrder
	orderItem           orderItem.BOrderItem
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger:              params.Logger,
		auth:                params.Auth,
		client:              params.Client,
		productInfo:         params.ProductInfo,
		product:             params.Product,
		productPriceHistory: params.ProductPriceHistory,
		status:              params.Status,
		trafficSource:       params.TrafficSource,
		order:               params.Order,
		orderItem:           params.OrderItem,
	}
}
