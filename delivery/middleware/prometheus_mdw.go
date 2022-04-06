package middleware

import (
	"enigmacamp.com/goprom/metric"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

func PrometheusUriRequestTotal() gin.HandlerFunc {
	return func(c *gin.Context) {
		metric.URIRequestTotal.With(prometheus.Labels{metric.Uri: c.Request.URL.RequestURI(), metric.Method: c.Request.Method}).Inc()
	}
}

func PrometheusUriErrorTotal() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/metrics" {
			c.Next()
			return
		}
		c.Next()
		status := strconv.Itoa(c.Writer.Status())
		if status[0] != '2' {
			println("Error")
			metric.URIErrorTotal.With(prometheus.Labels{metric.Uri: c.Request.URL.RequestURI(), metric.Method: c.Request.Method, metric.StatusCode: status}).Inc()
		}

	}
}

func PrometheusUriRequestDuration() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		fmt.Println("===>", latency.Seconds())
		metric.RequestDuration.With(prometheus.Labels{metric.Uri: c.Request.URL.RequestURI()}).Observe(latency.Seconds())
		metric.RequestDurationByTag.With(prometheus.Labels{"tag": c.GetString("tag")}).Observe(latency.Seconds())
	}
}
