package trafficsource

import (
	"github.com/ahliyor25/crm/internal/entities"
	trafficsource "github.com/ahliyor25/crm/internal/storage/trafficSource"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBTrafficSource)

// BTrafficSource ...
type BTrafficSource interface {
	Create(trafficSource entities.TrafficSource) (err error)
	Get(trafficSourceID uint) (data entities.TrafficSource, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	TrafficSource trafficsource.STrafficSource
	Logger        *logrus.Logger
}

// provider ...
type provider struct {
	trafficSource trafficsource.STrafficSource
	logger        *logrus.Logger
}

// NewBTrafficSource ...
func NewBTrafficSource(params Dependencies) BTrafficSource {
	return &provider{
		trafficSource: params.TrafficSource,
		logger:        params.Logger,
	}
}
