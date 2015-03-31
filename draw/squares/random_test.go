package squares

import (
	"image"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomGrid(t *testing.T) {
	t.Parallel()
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	RandomGrid(img, colorTheme, 10, float64(0.33))
}

func TestRandomGridSVG(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	RandomGridSVG(rec, colorTheme, 10, 10, 10, float64(50))
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}

func TestRandomGradientGrid(t *testing.T) {
	t.Parallel()
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	RandomGradientGrid(img, colorTheme, 10)

}

func TestRandomGradientGridSVG(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	RandomGradientGridSVG(rec, colorTheme, 10, 10, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}
