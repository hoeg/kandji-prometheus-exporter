package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Scraper struct {
	c *Collector
}

func StartHTTPS() {
	setupMetricsHandler()

	kandjiURL := os.Getenv("KANDJI_PROM_EXPORTER_KANDJI_URL")
	apiTokenFile := os.Getenv("KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE")
	token, err := loadAPITokenFromFile(apiTokenFile)
	if err != nil {
		fmt.Printf("Error loading API token: %s\n", err)
		return
	}
	port := os.Getenv("KANDJI_PROM_EXPORTER_PORT")

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
		log.Printf("error during scrape: %v", err)
		http.Error(w, "Error during scrape", http.StatusInternalServerError)
		return
	}
	log.Printf("scraped %d devices", len(devices))

	report := accumulateVersions(devices)
	for k, v := range report {
		versions.WithLabelValues(k).Set(float64(v))
	}
	log.Printf("Scrape successful!")
	w.WriteHeader(http.StatusOK)
}

func accumulateVersions(devices []Device) map[string]int {
	r := make(map[string]int)
	for _, device := range devices {
		os := fmt.Sprintf("%s %s", device.Platform, device.OsVersion)
		r[os] = r[os] + 1
	}
	return r
}

func loadAPITokenFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(content)), nil
}
