package squares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomGridSVG(t *testing.T) {
	rec := httptest.NewRecorder()
	RandomGridSVG(rec, colorTheme, 10, 10, 10, float64(50))
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}

func TestRandomGradientGridSVG(t *testing.T) {
	rec := httptest.NewRecorder()
	RandomGradientGridSVG(rec, colorTheme, 10, 10, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}
