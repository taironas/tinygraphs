package checkerboard

import (
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"image/color"
	"log"
	"net/http"
	"strconv"
)

// Color is the handler for /checkerboard/[0-8]
// build a 6x6 checkerboard with alternate colors based on the number passed in the url
func Color(w http.ResponseWriter, r *http.Request) {
	intID, err := misc.PermalinkID(r, 2)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := size(r)
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		colorMap := colors.MapOfColorPatterns()
		draw.Grid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		write.ImageJPEG(w, &img)
	}
}

// Checkerboard is the handler for /checkerboard/
// build a 6x6 checkerboard with alternate black and white colors.
func Checkerboard(w http.ResponseWriter, r *http.Request) {
	size := size(r)
	m := image.NewRGBA(image.Rect(0, 0, size, size))
	color1 := color.RGBA{uint8(255), uint8(255), 255, 255}
	color2 := color.RGBA{uint8(0), uint8(0), 0, 255}
	draw.Grid6X6(m, color1, color2)
	var img image.Image = m
	write.ImageJPEG(w, &img)
}

// extract size from HTTP request and return it.
func size(r *http.Request) int {
	strSize := r.FormValue("size")
	if len(strSize) > 0 {
		if size, errSize := strconv.ParseInt(strSize, 0, 64); errSize == nil {
			isize := int(size)
			if isize > 0 && isize < 1000 {
				return int(size)
			}
		}
	}
	return 210
}
