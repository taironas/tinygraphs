package spaceinvaders

import (
	"image/color"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

func SpaceInvaders(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
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
					colorMap[xQ] = draw.PickColor(key, colors, xQ+3*yQ)
				} else if xQ < squares {
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorMap[xQ] = colorMap[0]
				}
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, draw.FillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}
