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
	router.GET("/client", h.HClientGet, mw.JWT, mw.CORS)

	// ProductInfo
	router.POST("/productInfo", h.HProductInfoCreate, mw.JWT, mw.CORS)
	router.GET("/productInfo", h.HProductInfoGet, mw.JWT, mw.CORS)
	router.GET("/productInfo/getList", h.HProductInfoGetList, mw.JWT, mw.CORS)

	// Product
	router.POST("/product", h.HProductCreate, mw.JWT, mw.CORS)
	router.GET("/product", h.HProductGet, mw.JWT, mw.CORS)

	// Order
	router.POST("/order", h.HOrderCreate, mw.JWT, mw.CORS)
	router.GET("/order", h.HOrderGet, mw.JWT, mw.CORS)
	router.PUT("/order", h.HOrderUpdate, mw.JWT, mw.CORS)

	// Status
	router.POST("/status", h.HStatusCreate, mw.JWT, mw.CORS)
	router.GET("/status/{id}", h.HStatusGet, mw.JWT, mw.CORS)

	// TrafficSource
	router.POST("/trafficSource", h.HTrafficSourceCreate, mw.JWT, mw.CORS)
	router.GET("/trafficSource/{id}", h.HTrafficSourceGet, mw.JWT, mw.CORS)

	return
}
