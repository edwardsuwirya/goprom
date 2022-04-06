package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	Uri        = "uri"
	Method     = "method"
	StatusCode = "statusCode"
	Stats      = "stats"
)

var (
	URIRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goprom_uri_request_total",
			Help: "all the server received request num with every uri.",
		}, []string{Uri, Method},
	)

	URIErrorTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goprom_uri_error_total",
			Help: "all the server error request with every uri.",
		}, []string{Uri, Method, StatusCode},
	)

	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "goprom_request_duration",
			Help:    "the time server took to handle the request.",
			Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
		}, []string{Uri},
	)

	RequestDurationByTag = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "goprom_request_duration_tag",
			Help:    "the time server took to handle the request.",
			Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
		}, []string{"tag"},
	)

	MemoryUsage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goprom_memory_request",
			Help: "the memory took to run the app.",
		}, []string{Stats},
	)
)
