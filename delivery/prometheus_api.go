package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusApi struct {
	router *gin.Engine
}

func NewPrometheusApi(router *gin.Engine) *PrometheusApi {
	promApi := new(PrometheusApi)
	promApi.router = router
	promApi.initRouter()
	return promApi
}
func (api *PrometheusApi) initRouter() {
	api.router.GET("/metrics", func(context *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(context.Writer, context.Request)
	})
}
