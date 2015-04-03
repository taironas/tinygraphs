package isogrids

import (
	"image/color"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
)

// RandomGradientColor builds a isogrid image with with x colors selected at random for each quadrant.
// the background color stays the same the other colors get mixed in a gradient color from the first one to the last one.
func RandomGradientColor(w http.ResponseWriter, colors, gColors []color.RGBA, gv colors.GradientVector, width, height, lines int, prob float64) {

	var gradientColors []svg.Offcolor
	gradientColors = make([]svg.Offcolor, len(gColors))
	percentage := uint8(100 / len(gColors))

	step := uint8(100 / len(gColors))
	for i, c := range gColors {
		gradientColors[i] = svg.Offcolor{
			Offset:  percentage,
			Color:   draw.RGBToHex(c.R, c.G, c.B),
			Opacity: 1,
		}
		percentage += step
	}

	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Def()
	canvas.LinearGradient("gradientColors", gv.X1, gv.Y1, gv.X2, gv.Y2, gradientColors)
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill:url(#gradientColors)")

	fringeSize := width / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = width / fringeSize

	colorMap := make(map[int]color.RGBA)
	colorIndex := make(map[int]int)

	for xL := 0; xL <= lines; xL++ {
		colorMap = make(map[int]color.RGBA)
		colorIndex = make(map[int]int)
		for yL := -1; yL <= lines; yL++ {
			var x1, x2, y1, y2, y3 int
			var fill string
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}

			colorIndex[yL] = draw.RandomIndexFromArrayWithFreq(colors, prob)
			colorMap[yL] = colors[colorIndex[yL]]

			if colorIndex[yL] != 0 {
				fill = "fill:none"
			} else {
				fill = draw.FillFromRGBA(colorMap[yL])
			}

			canvas.Polygon(xs, ys, fill)

			var x11, x12, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11, y11, x12, y12, _, y13 = left2ndTriangle(xL, yL, fringeSize, distance)

				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y3
			} else {
				x11, y11, x12, y12, _, y13 = right2ndTriangle(xL, yL, fringeSize, distance)

				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y1 + fringeSize
			}
			xs1 := []int{x12, x11, x12}
			ys1 := []int{y11, y12, y13}

			colorIndex[yL] = draw.RandomIndexFromArrayWithFreq(colors, prob)
			colorMap[yL] = colors[colorIndex[yL]]

			if colorIndex[yL] != 0 {
				fill = "fill:none"
			} else {
				fill = draw.FillFromRGBA(colorMap[yL])
			}

			canvas.Polygon(xs1, ys1, fill)
		}
	}
	canvas.End()
}
