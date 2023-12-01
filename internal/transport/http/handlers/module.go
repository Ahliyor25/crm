package handlers

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger *logrus.Logger
}

// Handler ...
type Handler struct {
	logger *logrus.Logger
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger: params.Logger,
	}
}
