package scrape

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hoeg/kandji-prometheus-exporter/internal/collector"
)

type Scraper struct {
	c *collector.Collector
}

func StartHTTPS() {
	c, err := newConfig()
	if err != nil {
		log.Fatal("Invalid configuration: %+v", err)
	}
	s := Scraper{
		c: collector.New(c.kandjiURL, c.kandjiAPIToken),
	}
	setupMetricsHandler(c.metricsAPIToken)
	http.HandleFunc("/scrape", func(w http.ResponseWriter, r *http.Request) {
		s.scrapeHandler(w, r)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", c.port), nil)
}

func (s *Scraper) scrapeHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := s.c.ListDevices()
	if err != nil {
		log.Printf("error during scrape: %v", err)
		http.Error(w, "Error during scrape", http.StatusInternalServerError)
		return
	}
	log.Printf("scraped %d devices", len(devices))

	versionReport := collector.AccumulateVersions(devices)
	for k, v := range versionReport {
		versions.WithLabelValues(k).Set(float64(v))
	}

	blueprintReport := collector.Blueprints(devices)
	for k, v := range blueprintReport {
		blueprints.WithLabelValues(k).Set(float64(v))
	}

	log.Printf("Scrape successful!")
	w.WriteHeader(http.StatusOK)
}
