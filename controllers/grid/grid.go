package grid

import (
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"image/color"
	"log"
	"net/http"
)

// gridColorHandler is the handler for /grid/[0-8]
// build a 6x6 grid with alternate colors based on the number passed in the url
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
		write.Image(w, &img)
	}
}

// grid6X6Handler is the handler for /grid/
// build a 6x6 grid with alternate black and white colors.
func H6X6(w http.ResponseWriter, r *http.Request) {
	size := size(r)
	m := image.NewRGBA(image.Rect(0, 0, size, size))
	color1 := color.RGBA{uint8(255), uint8(255), 255, 255}
	color2 := color.RGBA{uint8(0), uint8(0), 0, 255}
	draw.Grid6X6(m, color1, color2)
	var img image.Image = m
	write.Image(w, &img)
}
