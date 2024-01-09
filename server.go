package main

import (
	"net/http"
)

func startHTTPServer() {
	setupMetricsHandler()

	http.HandleFunc("/scrape", func(w http.ResponseWriter, r *http.Request) {
		scrapeHandler(w, r)
	})

	http.ListenAndServe(":8080", nil)
}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	err := makeHTTPSRequest()

	if err != nil {
		errors.Inc()
		http.Error(w, "Error during scrape", http.StatusInternalServerError)
		return
	}

	successfulRequests.Inc()
	fmt.Fprint(w, "Scrape successful!")
}

