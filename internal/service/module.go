package service

import (
	"github.com/ahliyor25/crm/internal/service/auth"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	auth.Module,
)
