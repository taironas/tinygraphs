package squares

import (
	"image"
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// handler for "/squares/random"
// generates a black and white grid random image.
func Random(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	theme := extract.Theme(r)
	colorMap := colors.MapOfColorThemes()

	var bg, fg color.RGBA
	var err error

	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	} else {
		if bg, err = extract.Background(r); err != nil {
			bg = colorMap["base"][0]
		}
		if fg, err = extract.Foreground(r); err != nil {
			fg = colorMap["base"][1]
		}
	}

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.RandomGrid6X6(m, bg, fg)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		draw.RandomGrid6X6SVG(w, bg, fg, size)
	}
}
