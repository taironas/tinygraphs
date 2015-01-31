package checkerboard

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckerboard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Checkerboard))
	defer ts.Close()

	req, err := http.NewRequest("GET", "http://tinygraphs.com/checkerboard", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	Checkerboard(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Page didn't return %v", http.StatusOK)
	}
}
