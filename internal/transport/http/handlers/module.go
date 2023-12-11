package handlers

import (
	"github.com/ahliyor25/crm/internal/service/auth"
	"github.com/ahliyor25/crm/internal/service/client"
	"github.com/ahliyor25/crm/internal/service/product"
	productinfo "github.com/ahliyor25/crm/internal/service/productInfo"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger      *logrus.Logger
	Auth        auth.BAuth
	Client      client.BClient
	ProductInfo productinfo.BProductInfo
	Product     product.BProduct
}

// Handler ...
type Handler struct {
	logger      *logrus.Logger
	auth        auth.BAuth
	client      client.BClient
	productInfo productinfo.BProductInfo
	product     product.BProduct
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger:      params.Logger,
		auth:        params.Auth,
		client:      params.Client,
		productInfo: params.ProductInfo,
		product:     params.Product,
	}
}
