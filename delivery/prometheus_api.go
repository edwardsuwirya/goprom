package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusApi struct {
	publicRoute *gin.RouterGroup
}

func NewPrometheusApi(publicRoute *gin.RouterGroup) *PrometheusApi {
	studentApi := new(PrometheusApi)
	studentApi.publicRoute = publicRoute
	studentApi.initRouter()
	return studentApi
}
func (api *PrometheusApi) initRouter() {
	userRoute := api.publicRoute.Group("/metrics")
	userRoute.GET("", func(context *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(context.Writer, context.Request)
	})
}
