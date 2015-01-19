package draw

import (
	"github.com/ajstarks/svgo"
	"image/color"
	"math"
	"net/http"
)

// IsogridsSVG builds an image with 10x10 grids of alternate colors.
func IsogridsSVG(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	squares := 6
	quadrantSize := size / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)

		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			if _, ok := colorMap[xQ]; !ok {
				if float64(xQ) < middle {
					colorMap[xQ] = colorFromKey(key, color1, color2, xQ+3*yQ)
				} else if xQ < squares {
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorMap[xQ] = colorMap[0]
				}
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, fillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}
