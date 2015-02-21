package squares

import (
	"image/color"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// GradientSVG builds an image with 6 by 6 quadrants of alternate colors.
func GradientSVG(w http.ResponseWriter, key string, colors []color.RGBA, size int) {

	rainbow := []svg.Offcolor{
		{10, "#00cc00", 1},
		{30, "#006600", 1},
		{70, "#cc0000", 1},
		{90, "#000099", 1}}
	canvas := svg.New(w)
	canvas.Start(size, size)
	canvas.Def()
	canvas.LinearGradient("rainbow", 0, 0, uint8(size), 0, rainbow)
	canvas.DefEnd()
	canvas.Rect(0, 0, size, size, "fill:url(#rainbow)")

	squares := 6
	quadrantSize := size / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	colorIndex := make(map[int]int)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)
		colorIndex = make(map[int]int)
		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			fill := ""
			if _, ok := colorMap[xQ]; !ok {
				if float64(xQ) < middle {
					colorIndex[xQ] = draw.PickIndex(key, len(colors), xQ+3*yQ)
					colorMap[xQ] = draw.PickColor(key, colors, xQ+3*yQ)
				} else if xQ < squares {
					colorIndex[xQ] = colorIndex[squares-xQ-1]
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorIndex[xQ] = colorIndex[0]
					colorMap[xQ] = colorMap[0]
				}
			}
			if colorIndex[xQ] != 0 {
				fill = "fill:none; opacity:0.3"
			} else {
				fill = draw.FillFromRGBA(colorMap[xQ])
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, fill)
		}
	}
	canvas.End()
}
