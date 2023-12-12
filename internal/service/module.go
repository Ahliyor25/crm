package service

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

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	auth.Module,
	client.Module,
	productinfo.Module,
	product.Module,
	productpricehistory.Module,
	order.Module,
	orderItem.Module,
	status.Module,
	trafficsource.Module,
)
