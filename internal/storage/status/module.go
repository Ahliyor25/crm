package status

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSStatus)

// SStatus ...
type SStatus interface {
	Create(Status entities.Status) (err error)
	Get(target entities.Status) (data entities.Status, err error)
	Update(data entities.Status) (err error)
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

// NewSStatus ...
func NewSStatus(params Dependencies) SStatus {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
