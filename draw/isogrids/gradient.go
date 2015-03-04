package isogrids

import (
	"image/color"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
)

// RandomGradientSVG builds an image.
func RandomGradientSVG(w http.ResponseWriter, colors, gColors []color.RGBA, gv colors.GradientVector, width, height, xsquares int) {

	var gradientColors []svg.Offcolor
	gradientColors = make([]svg.Offcolor, len(gColors))
	percentage := uint8(100 / len(gColors))

	step := uint8(100 / len(gColors))
	for i, c := range gColors {
		gradientColors[i] = svg.Offcolor{percentage, RGBToHex(c.R, c.G, c.B), 1}
		percentage += step
	}

	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Def()
	canvas.LinearGradient("gradientColors", gv.X1, gv.Y1, gv.X2, gv.Y2, gradientColors)
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill:url(#gradientColors)")

	squares := xsquares
	quadrantSize := width / squares
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
					colorIndex[xQ] = draw.RandomIndexFromArray(colors)
					colorMap[xQ] = colors[colorIndex[xQ]]
				} else if xQ < squares {
					colorIndex[xQ] = colorIndex[squares-xQ-1]
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorIndex[xQ] = colorIndex[0]
					colorMap[xQ] = colorMap[0]
				}
			}
			if colorIndex[xQ] != 0 {
				fill = "fill:none"
			} else {
				fill = draw.FillFromRGBA(colorMap[xQ])

			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, fill)
		}
	}
	canvas.End()
}
