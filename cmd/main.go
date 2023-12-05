package main

import (
	"github.com/ahliyor25/crm/internal/service"
	"github.com/ahliyor25/crm/internal/storage"
	"github.com/ahliyor25/crm/internal/transport/http/handlers"
	"github.com/ahliyor25/crm/internal/transport/http/router"

	"github.com/ahliyor25/crm/pkg/bootstrap/http"
	"github.com/ahliyor25/crm/pkg/config"
	"github.com/ahliyor25/crm/pkg/databases"
	"github.com/ahliyor25/crm/pkg/logger"
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
