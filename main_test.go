package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapeHandler(t *testing.T) {
	t.Skip()
	t.Run("Successful Scrape", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/scrape", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(scrapeHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := "Scrape successful!"
		if rr.Body.String() != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}

// Add more tests for error cases, HTTPS request, and other scenarios.
