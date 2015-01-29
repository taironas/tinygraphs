package checkerboard

import (
	"image"
	"image/color"
	"log"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// Checkerboard is the handler for /checkerboard
// build a 6x6 checkerboard with alternate black and white colors.
func Checkerboard(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	theme := extract.Theme(r)
	colorMap := colors.MapOfColorThemes()

	var c1, c2 color.RGBA
	if val, ok := colorMap[theme]; ok {
		c1 = val[0]
		c2 = val[1]
	} else {
		c1 = colorMap["base"][0]
		c2 = colorMap["base"][1]
	}
	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		squares.Grid6X6(m, c1, c2)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.Grid6X6SVG(w, c1, c2, size)
	}
}
