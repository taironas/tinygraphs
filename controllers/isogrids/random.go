package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Random handler for /isogrids/random.
// Generates a random isogrid image.
func Random(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorThemes()
	bg, fg := extract.ExtraColors(r, colorMap)
	theme := extract.Theme(r)
	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	}

	var colors []color.RGBA
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			numColors := extract.NumColors(r)
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}

	if newColors, err := extract.Colors(r); err == nil {
		colors = newColors
	}

	prob := extract.Probability(r, 1/float64(len(colors)))

	size := extract.Size(r)
	lines := extract.Lines(r)
	write.ImageSVG(w)
	isogrids.Random(w, "", colors, size, size, lines, prob)
}
