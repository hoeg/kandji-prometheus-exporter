package scrape

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type config struct {
	kandjiURL       string
	kandjiAPIToken  string
	metricsAPIToken string
	port            string
}

func newConfig() (*config, error) {
	kandjiURL := os.Getenv("KANDJI_PROM_EXPORTER_KANDJI_URL")
	if kandjiURL == "" {
		log.Fatal("KANDJI_PROM_EXPORTER_KANDJI_URL must be set")
	}

	kandjiAPIToken, err := loadAPITokenFromEnv("KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE", false)
	if err != nil {
		return nil, fmt.Errorf("error loading Kandji API token: %w", err)

	}

	metricsAPIToken, err := loadAPITokenFromEnv("KANDJI_PROM_EXPORTER_METRICS_API_KEY_FILE", true)
	if err != nil {
		return nil, fmt.Errorf("error loading metrics API token: %w", err)

	}
	port := port()
	return &config{
		kandjiURL:       kandjiURL,
		kandjiAPIToken:  kandjiAPIToken,
		metricsAPIToken: metricsAPIToken,
		port:            port,
	}, err
}

func port() string {
	port := os.Getenv("KANDJI_PROM_EXPORTER_PORT")
	if port == "" {
		port = "8080"
		log.Printf("No port set with KANDJI_PROM_EXPORTER_PORT defaulting to port %s", port)
	}
	return port
}

func loadAPITokenFromEnv(envPath string, optional bool) (string, error) {
	tokenFilePath := os.Getenv(envPath)
	if tokenFilePath == "" {
		if optional {
			return "", nil
		}
		return "", fmt.Errorf("expected %s to be set", envPath)
	}
	content, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(content)), nil
}
