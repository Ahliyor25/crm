package misc

import (
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	response.Module,
)
