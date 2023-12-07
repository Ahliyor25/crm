package user

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSLogin)

// SUser ...
type SUser interface {
	Create(user entities.User) (err error)
	Get(target entities.User) (data entities.User, err error)
	GetMany(target entities.User, currentUserID uint) (users []entities.User, err error)
	Update(data entities.User) (err error)
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
