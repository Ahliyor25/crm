package storage

import (
	"github.com/ahliyor25/crm/internal/storage/client"
	"github.com/ahliyor25/crm/internal/storage/order"
	"github.com/ahliyor25/crm/internal/storage/orderItem"
	"github.com/ahliyor25/crm/internal/storage/product"
	productinfo "github.com/ahliyor25/crm/internal/storage/productInfo"
	productpricehistory "github.com/ahliyor25/crm/internal/storage/productPriceHistory"
	"github.com/ahliyor25/crm/internal/storage/status"
	trafficsource "github.com/ahliyor25/crm/internal/storage/trafficSource"
	"github.com/ahliyor25/crm/internal/storage/user"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	client.Module,
	user.Module,
	productinfo.Module,
	product.Module,
	productpricehistory.Module,
	status.Module,
	trafficsource.Module,
	order.Module,
	orderItem.Module,
)
