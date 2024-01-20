package scrape

import (
	"crypto/subtle"
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
		[]string{"blueprint_name"})
)

func init() {
	prometheus.MustRegister(versions, blueprints)
}

func setupMetricsHandler(metricsAPIToken string) {
	var metricsHandler http.Handler
	if metricsAPIToken == "" {
		metricsHandler = promhttp.Handler()
	} else {
		auth := auth{
			apiToken: metricsAPIToken
		}
		metricsHandler = auth.middleware(promhttp.Handler())
	}
	http.Handle("/metrics", metricsHandler)
}

type auth struct {
	apiToken string
}

func (a *auth) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if subtle.ConstantTimeCompare([]byte(authHeader), []byte(a.apiToken)) != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
