package app

import "github.com/prometheus/client_golang/prometheus"

var (
	customCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_custom_request_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "status"},
	)

	customErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_custom_error_total",
			Help: "Total number of errors returned by the API",
		},
		[]string{"path", "status"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_custom_request_duration_seconds",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2, 5, 7, 10},
		},
		[]string{"response_status"},
	)
)

func init() {
	prometheus.MustRegister(
		customCounter,
		customErrorCounter,
		requestDuration,
	)
}
