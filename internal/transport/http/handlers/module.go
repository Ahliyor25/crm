package handlers

import (
	"github.com/ahliyor25/crm/internal/service/auth"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger *logrus.Logger
	Auth   auth.BAuth
}

// Handler ...
type Handler struct {
	logger *logrus.Logger
	auth   auth.BAuth
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger: params.Logger,
		auth:   params.Auth,
	}
}
