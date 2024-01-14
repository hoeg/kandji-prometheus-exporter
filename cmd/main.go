package main

import "github.com/hoeg/kandji-prometheus-exporter/internal/scrape"

func main() {
	scrape.StartHTTPS()
}
