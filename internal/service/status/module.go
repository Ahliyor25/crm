package status

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/internal/storage/status"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBStatus)

// BStatus ...
type BStatus interface {
	Create(status entities.Status) (err error)
	Get(statusID uint) (data entities.Status, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Status status.SStatus
	Logger *logrus.Logger
}

// provider ...
type provider struct {
	status status.SStatus
	logger *logrus.Logger
}

// NewBStatus ...
func NewBStatus(params Dependencies) BStatus {
	return &provider{
		status: params.Status,
		logger: params.Logger,
	}
}
