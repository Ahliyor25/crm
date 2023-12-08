package client

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSClient)

// SClient ...
type SClient interface {
	Create(client entities.Client) (err error)
	Get(target entities.Client) (data entities.Client, err error)
	GetMany(target entities.Client, currentClientID uint) (clients []entities.Client, err error)
	Update(data entities.Client) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

// provider ...
type provider struct {
	logger   *logrus.Logger
	postgres *gorm.DB
}

// NewSClient ...
func NewSClient(params Dependencies) SClient {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
