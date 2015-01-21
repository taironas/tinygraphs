package checkerboard

import (
	"image"
	"image/color"
	"log"
	"net/http"
	"strconv"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// Color is the handler for /checkerboard/:colorId
// build a 6x6 checkerboard with alternate colors based on the number passed in the url
func Color(w http.ResponseWriter, r *http.Request) {
	id := route.Context.Get(r, "colorId")
	if colorId, err := strconv.ParseInt(id, 0, 64); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := extract.Size(r)
		colorMap := colors.MapOfColorPatterns()
		var c1, c2 color.RGBA
		if val, ok := colorMap[int(colorId)]; ok {
			c1 = val[0]
			c2 = val[1]
		} else {
			c1 = colorMap[0][0]
			c2 = colorMap[0][1]
		}

		if f := extract.Format(r); f == format.JPEG {
			m := image.NewRGBA(image.Rect(0, 0, size, size))
			draw.Grid6X6(m, c1, c2)
			var img image.Image = m
			write.ImageJPEG(w, &img)
		} else if f == format.SVG {
			write.ImageSVG(w)
			draw.Grid6X6SVG(w, c1, c2, size)
		}
	}
}

// Checkerboard is the handler for /checkerboard
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
		write.ImageSVG(w)
		draw.Grid6X6SVG(w, color1, color2, size)
	}
}
