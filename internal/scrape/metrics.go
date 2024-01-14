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
	blueprints = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kandji_device_blueprint",
			Help: "Number of devices on the different blueprints",
		},
		[]string{"blueprint_name"}
	)
)

func init() {
	prometheus.MustRegister(versions, blueprints)
}

func setupMetricsHandler() {
	http.Handle("/metrics", promhttp.Handler())
}
