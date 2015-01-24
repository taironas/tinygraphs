package squares

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

// handler for "/squares/random"
// generates a black and white grid random image.
func Random(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	colorMap := colors.MapOfColorPatterns()
	bg, err1 := extract.Background(r)
	if err1 != nil {
		bg = colorMap[0][0]
	}
	fg, err2 := extract.Foreground(r)
	if err2 != nil {
		fg = colorMap[0][1]
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

// handler for "/squares/random/:colorId"
// generates a grid random image with a specific color based on the colorMap
func RandomColor(w http.ResponseWriter, r *http.Request) {
	id, _ := route.Context.Get(r, "colorId")
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
			draw.RandomGrid6X6(m, c1, c2)
			var img image.Image = m
			write.ImageJPEG(w, &img)
		} else if f == format.SVG {
			write.ImageSVG(w)
			draw.RandomGrid6X6SVG(w, c1, c2, size)
		}
	}
}
