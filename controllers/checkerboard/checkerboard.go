package checkerboard

import (
	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"image/color"
	"log"
	"net/http"
)

// Color is the handler for /checkerboard/[0-8]
// build a 6x6 checkerboard with alternate colors based on the number passed in the url
func Color(w http.ResponseWriter, r *http.Request) {
	intID, err := misc.PermalinkID(r, 2)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := extract.Size(r)
		colorMap := colors.MapOfColorPatterns()
		if f := extract.Format(r); f == format.JPEG {
			m := image.NewRGBA(image.Rect(0, 0, size, size))
			draw.Grid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
			var img image.Image = m
			write.ImageJPEG(w, &img)
		} else if f == format.SVG {
			canvas := svg.New(w)
			draw.Grid6X6SVG(canvas, colorMap[int(intID)][0], colorMap[int(intID)][1], size)
			write.ImageSVG(w, canvas)
		}
	}
}

// Checkerboard is the handler for /checkerboard/
// build a 6x6 checkerboard with alternate black and white colors.
func Checkerboard(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	color1 := color.RGBA{uint8(255), uint8(255), 255, 255}
	color2 := color.RGBA{uint8(0), uint8(0), 0, 255}
	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.Grid6X6(m, color1, color2)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		canvas := svg.New(w)
		draw.Grid6X6SVG(canvas, color1, color2, size)
		write.ImageSVG(w, canvas)
	}
}
