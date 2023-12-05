package user

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSLogin)

// SUser ...
type SUser interface {
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

type provider struct {
	logger   *logrus.Logger
	postgres *gorm.DB
}

// NewSLogin ...
func NewSLogin(params Dependencies) SUser {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
