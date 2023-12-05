package server

import (
	"fmt"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/router"
	"github.com/ahliyor25/crm/pkg/config"
	"go.uber.org/fx"
	"net"
	"net/http"
)

// ServerModule ...
var ServerModule = fx.Provide(NewServer)

// Dependecies ...
type Dependecies struct {
	fx.In

	Config *config.Config
	Router *router.HTTPRouter
}

// NewServer ...
func NewServer(params Dependecies) *http.Server {
	url := net.JoinHostPort(params.Config.Server.Host, fmt.Sprint(params.Config.Server.Port))

	return &http.Server{
		MaxHeaderBytes: 32 << 20, // 32 Mb
		Addr:           url,
		Handler:        params.Router,
	}
}
