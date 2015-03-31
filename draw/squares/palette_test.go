package squares

import (
	"image"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPalette(t *testing.T) {
	t.Parallel()
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	Palette(img, colorTheme)
}

func TestPaletteSVG(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	PaletteSVG(rec, colorTheme, 100, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}
