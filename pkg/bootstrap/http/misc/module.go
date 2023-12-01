package misc

import (
	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/misc/response"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	response.Module,
)
