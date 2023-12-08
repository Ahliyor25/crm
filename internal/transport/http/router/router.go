package router

import (
	"github.com/ahliyor25/crm/internal/transport/http/handlers"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/middlewares"
	transportHTTP "github.com/ahliyor25/crm/pkg/bootstrap/http/router"
)

// NewRouter ..
func NewRouter(h *handlers.Handler, mw middlewares.Middleware) (router *transportHTTP.HTTPRouter) {
	router = transportHTTP.NewRouter()
	router.ConnectSwagger(h.ServeSwaggerFiles)
	router.GET("/ping", h.HPingPong, mw.JWT, mw.CORS)
	router.POST("/auth/login", h.HAuthLogin, mw.CORS)

	// Client
	router.POST("/client", h.HClientCreate, mw.JWT, mw.CORS)

	return
}
