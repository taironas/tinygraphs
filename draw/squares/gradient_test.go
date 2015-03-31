package squares

import (
	"image/color"
	"net/http"
	"net/http/httptest"
	"testing"

	tgColors "github.com/taironas/tinygraphs/colors"
)

var (
	colorTheme = []color.RGBA{
		color.RGBA{255, 245, 249, 255},
		color.RGBA{232, 70, 134, 255},
		color.RGBA{232, 70, 186, 255},
		color.RGBA{232, 70, 81, 255},
	}
)

func TestRandomGradientColorSVG(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	gv := tgColors.GradientVector{0, 0, 1, 1}
	RandomGradientColorSVG(rec, colorTheme, colorTheme[1:], gv, 10, 10, 10, float64(50))
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}
