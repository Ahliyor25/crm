package http

import (
	"github.com/ahliyor25/crm/pkg/bootstrap/http/middlewares"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/server"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ModuleLifecycleHooks,
	server.ServerModule,
)
