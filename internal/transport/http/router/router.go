package router

import (
	"github.com/waffle-frame/clean-architecture-template/internal/transport/http/handlers"
	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/middlewares"
	transportHTTP "github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/router"
)

// NewRouter ..
func NewRouter(h *handlers.Handler, mw middlewares.Middleware) (router *transportHTTP.HTTPRouter) {
	router = transportHTTP.NewRouter()
	router.ConnectSwagger(h.ServeSwaggerFiles)

	router.GET("/ping", h.HPingPong, mw.CORS)

	return
}
