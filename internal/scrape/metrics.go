package scrape

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	versions = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kandji_mac_os_version",
			Help: "Number of machines on different MacOS versions",
		},
		[]string{"os_version"})
)

func init() {
	prometheus.MustRegister(versions)
}

func setupMetricsHandler() {
	http.Handle("/metrics", promhttp.Handler())
}
