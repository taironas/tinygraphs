package grid

import (
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"

	"image"
	"log"
	"net/http"
)

// handler for "/gird/random/symetric/x"
// generates a black and white grid random image.
func RandomSymetricX(w http.ResponseWriter, r *http.Request) {
	size := size(r)
	colorMap := colors.MapOfColorPatterns()
	bg, err1 := background(r)
	if err1 != nil {
		bg = colorMap[0][0]
	}
	fg, err2 := foreground(r)
	if err2 != nil {
		fg = colorMap[0][1]
	}
	m := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.RandomSymetricInXGrid6X6(m, bg, fg)
	var img image.Image = m
	write.Image(w, &img)
}

// handler for "/gird/random/symetric/y"
// generates a black and white grid random image.
func RandomSymetricY(w http.ResponseWriter, r *http.Request) {
	size := size(r)
	colorMap := colors.MapOfColorPatterns()
	bg, err1 := background(r)
	if err1 != nil {
		bg = colorMap[0][0]
	}
	fg, err2 := foreground(r)
	if err2 != nil {
		fg = colorMap[0][1]
	}
	m := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.RandomSymetricInYGrid6X6(m, bg, fg)
	var img image.Image = m
	write.Image(w, &img)
}

// handler for "/grid/random/symetric/y/[0-9]"
// generates a grid random image with a specific color based on the colorMap
func RandomSymetricYColor(w http.ResponseWriter, r *http.Request) {
	intID, err := misc.PermalinkID(r, 5)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := size(r)
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		colorMap := colors.MapOfColorPatterns()
		draw.RandomSymetricInYGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		write.Image(w, &img)
	}
}

// handler for "/grid/random/symetric/x/[0-9]"
// generates a grid random image with a specific color based on the colorMap
func RandomSymetricXColor(w http.ResponseWriter, r *http.Request) {
	intID, err := misc.PermalinkID(r, 5)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := size(r)
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		colorMap := colors.MapOfColorPatterns()
		draw.RandomSymetricInXGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		write.Image(w, &img)
	}
}
