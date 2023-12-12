package trafficsource

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewSTrafficSource)

// STrafficSource ...
type STrafficSource interface {
	Create(trafficSource entities.TrafficSource) (err error)
	Get(target entities.TrafficSource) (data entities.TrafficSource, err error)
	Update(data entities.TrafficSource) (err error)
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

// NewSTrafficSource ...
func NewSTrafficSource(params Dependencies) STrafficSource {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
