package gradient

import (
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/write"

	"image"
	"net/http"
)

// Gradient is the handler for /gradient/
func Gradient(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	draw.Gradient(m)
	var img image.Image = m
	write.Image(w, &img)
}
