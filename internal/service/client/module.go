package client

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/internal/storage/client"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBCLient)

// BClient ...
type BClient interface {
	Create(client entities.Client) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Client client.SClient
	Logger *logrus.Logger
}

// provider ...
type provider struct {
	client client.SClient
	logger *logrus.Logger
}

// NewBCLient ...
func NewBCLient(params Dependencies) BClient {
	return &provider{
		client: params.Client,
		logger: params.Logger,
	}
}
