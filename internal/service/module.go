package service

import (
	"github.com/ahliyor25/crm/internal/service/auth"
	"github.com/ahliyor25/crm/internal/service/client"
	"github.com/ahliyor25/crm/internal/service/product"
	productinfo "github.com/ahliyor25/crm/internal/service/productInfo"
	
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	auth.Module,
	client.Module,
	productinfo.Module,
	product.Module,
)
