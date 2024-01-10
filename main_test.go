package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapeHandler(t *testing.T) {
	t.Run("Successful Scrape", func(t *testing.T) {
		nothingLeft := false
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if nothingLeft {
				fmt.Fprint(w, "")
				return
			}
			devices := []Device{}
			devices = append(devices, Device{Platform: "Mac", OsVersion: "1"})

			out, err := json.Marshal(devices)
			if err != nil {
				fmt.Fprintf(w, "error: %s", err)
			}
			w.Write(out)
			nothingLeft = true
		}))
		defer mockServer.Close()

		// Replace myPlaceURL with the mock server URL
		s := Scraper{
			c: NewCollector(mockServer.URL, "token"),
		}

		req, err := http.NewRequest("GET", "/scrape", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.scrapeHandler)

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

// Add more tests for error cases, HTTPS request, and other scenarios
