package main

import (
	"github.com/waffle-frame/clean-architecture-template/internal/service"
	"github.com/waffle-frame/clean-architecture-template/internal/storage"
	"github.com/waffle-frame/clean-architecture-template/internal/transport/http/handlers"
	"github.com/waffle-frame/clean-architecture-template/internal/transport/http/router"

	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http"
	"github.com/waffle-frame/clean-architecture-template/pkg/config"
	"github.com/waffle-frame/clean-architecture-template/pkg/databases"
	"github.com/waffle-frame/clean-architecture-template/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		config.Module,
		logger.Module,
		databases.PostgresModule,

		service.Module,
		storage.Module,
		handlers.Module,
		router.Module,
	)

	app.Run()
}
