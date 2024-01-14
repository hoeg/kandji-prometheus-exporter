package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
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
			t.Fatalf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		if float64(1) != testutil.ToFloat64(versions.WithLabelValues("Mac 1")) {
			t.Fatal("expected 1 instance of Mac 1")
		}
	})
}

// Add more tests for error cases, HTTPS request, and other scenarios
