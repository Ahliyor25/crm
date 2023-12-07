package storage

import (
	"github.com/ahliyor25/crm/internal/storage/user"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	user.Module,
)
