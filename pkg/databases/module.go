package databases

import (
	"github.com/waffle-frame/clean-architecture-template/pkg/config"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// PostgresModule ...
var PostgresModule = fx.Provide(NewPostgresConn)

// Dependencies ...
type Dependencies struct {
	fx.In

	Logger *logrus.Logger
	Config *config.Config
}

// NewPostgresConn ...
func NewPostgresConn(params Dependencies) *gorm.DB {
	return Postgres(params)
}
