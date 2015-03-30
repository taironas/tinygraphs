package isogrids

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandom(t *testing.T) {

	rec := httptest.NewRecorder()
	Random(rec, key, colorTheme, 10, 10, 10, float64(50))
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}

func TestRandomGradient(t *testing.T) {

	rec := httptest.NewRecorder()
	RandomGradient(rec, key, colorTheme, 10, 10, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}

func TestRandomMirror(t *testing.T) {

	rec := httptest.NewRecorder()
	RandomMirror(rec, key, colorTheme, 10, float64(50))
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}
