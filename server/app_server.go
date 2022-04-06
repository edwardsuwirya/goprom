package server

import (
	"enigmacamp.com/goprom/config"
	"enigmacamp.com/goprom/delivery"
	"enigmacamp.com/goprom/delivery/middleware"
	"enigmacamp.com/goprom/metric"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	apiBaseUrl   string
	apiGroup     string
}

func (a *appServer) instrumentationMiddleware() {
	metric.ExecuteCollector(metric.NewMemoryUsageCollector(metric.MemoryUsage))
	a.routerEngine.Use(middleware.PrometheusUriRequestTotal())
	a.routerEngine.Use(middleware.PrometheusUriRequestDuration())
	a.routerEngine.Use(middleware.PrometheusUriErrorTotal())
}

func (a *appServer) handlers() {
	publicRoute := a.routerEngine.Group(a.apiGroup)
	a.routerEngine.GET("/metrics", func(context *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(context.Writer, context.Request)
	})
	delivery.NewStudentApi(publicRoute)
}

func (a *appServer) Run() {
	a.instrumentationMiddleware()
	a.handlers()
	err := a.routerEngine.Run(a.apiBaseUrl)
	if err != nil {
		panic("Error run server")
	}
}

func NewAppServer() AppServer {
	newServer := new(appServer)

	c := config.NewConfig()

	r := gin.Default()
	newServer.routerEngine = r
	newServer.apiBaseUrl = c.ApiBaseUrl()
	newServer.apiGroup = c.ApiGroup()
	return newServer
}
