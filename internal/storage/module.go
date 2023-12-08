package storage

import (
	"github.com/ahliyor25/crm/internal/storage/client"
	"github.com/ahliyor25/crm/internal/storage/product"
	productinfo "github.com/ahliyor25/crm/internal/storage/productInfo"
	"github.com/ahliyor25/crm/internal/storage/status"
	trafficsource "github.com/ahliyor25/crm/internal/storage/trafficSource"
	"github.com/ahliyor25/crm/internal/storage/user"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	user.Module,
	client.Module,
	trafficsource.Module,
	status.Module,
	productinfo.Module,
	product.Module,
)
