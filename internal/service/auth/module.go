package auth

import (
	"github.com/ahliyor25/crm/internal/storage/user"
	"github.com/ahliyor25/crm/pkg/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewBAuth)

// BAuth ...
type BAuth interface {
	Login(email, password string) (token, refreshToken string, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger
	User   user.SUser
	Config *config.Config
}

// provider ...
type provider struct {
	logger *logrus.Logger
	user   user.SUser
	config *config.Config
}

// NewBAuth ...
func NewBAuth(params Dependencies) BAuth {
	return &provider{
		logger: params.Logger,
		user:   params.User,
		config: params.Config,
	}
}
