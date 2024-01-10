package server

import (
	"fmt"
	"net/http"
)

type Scraper struct {
	c *Collector
}

func StartHTTPS() {
	setupMetricsHandler()

	//read config
	kandjiURL := ""
	token := ""
	port := ""

	s := Scraper{
		c: NewCollector(kandjiURL, token),
	}

	http.HandleFunc("/scrape", func(w http.ResponseWriter, r *http.Request) {
		s.scrapeHandler(w, r)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func (s *Scraper) scrapeHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := s.c.ListDevices()
	if err != nil {
		http.Error(w, "Error during scrape", http.StatusInternalServerError)
		return
	}

	report := accumulateVersions(devices)
	for k, v := range report {
		versions.WithLabelValues(k).Set(float64(v))
	}

	fmt.Fprint(w, "Scrape successful!")
}

func accumulateVersions(devices []Device) map[string]int {
	r := make(map[string]int)
	for _, device := range devices {
		os := fmt.Sprintf("%s %s", device.Platform, device.OsVersion)
		r[os] = r[os] + 1
	}
	return r
}
