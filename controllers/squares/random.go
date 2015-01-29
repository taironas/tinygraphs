package squares

import (
	"image"
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// Random handler for "/squares/random"
// generates a random 6 by 6 grid image.
func Random(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	numColors := extract.NumColors(r)

	colorMap := colors.MapOfColorThemes()

	bg, fg := extract.ExtraColors(r, colorMap)

	var colors []color.RGBA
	theme := extract.Theme(r)
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		squares.RandomGrid6X6(m, colors)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.RandomGrid6X6SVG(w, colors, size)
	}
}
