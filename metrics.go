package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	successfulRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myplace_successful_requests_total",
		Help: "Total number of successful requests to myplace.org",
	})
	errors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myplace_errors_total",
		Help: "Total number of errors during requests to myplace.org",
	})
)

func init() {
	prometheus.MustRegister(successfulRequests, errors)
}

func setupMetricsHandler() {
	http.Handle("/metrics", promhttp.Handler())
}
