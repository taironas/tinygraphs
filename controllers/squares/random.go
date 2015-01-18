package squares

import (
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"log"
	"net/http"
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

// handler for "/squares/random/[0-9]"
// generates a grid random image with a specific color based on the colorMap
func RandomColor(w http.ResponseWriter, r *http.Request) {
	intID, err := misc.PermalinkID(r, 3)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := extract.Size(r)
		colorMap := colors.MapOfColorPatterns()
		if f := extract.Format(r); f == format.JPEG {
			m := image.NewRGBA(image.Rect(0, 0, size, size))
			draw.RandomGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
			var img image.Image = m
			write.ImageJPEG(w, &img)
		} else if f == format.SVG {
			write.ImageSVG(w)
			draw.RandomGrid6X6SVG(w, colorMap[int(intID)][0], colorMap[int(intID)][1], size)

		}
	}
}
