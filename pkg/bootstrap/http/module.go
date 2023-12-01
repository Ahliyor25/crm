package http

import (
	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/middlewares"
	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/misc"
	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/server"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ModuleLifecycleHooks,
	server.ServerModule,
)
