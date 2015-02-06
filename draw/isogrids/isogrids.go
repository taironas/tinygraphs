package isogrids

import (
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// Isogrids builds an image with 10x10 grids of half diagonals
func Isogrids(w http.ResponseWriter, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = size / fringeSize

	// triangle grid here:
	for xL := 0; xL < lines/2; xL++ {
		for yL := 0; yL < lines; yL++ {
			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}
			fill1 := draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+lines)%15))
			canvas.Polygon(xs, ys, fill1)

			xsMirror := mirrorCoordinates(xs, lines, fringeSize, 0)
			canvas.Polygon(xsMirror, ys, fill1)

			var x11, x12, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11, y11, x12, y12, _, y13 = left2ndTriangle(xL, yL, fringeSize, distance)

				// in order to have a perfect hexagon,
				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y3
			} else {
				x11, y11, x12, y12, _, y13 = right2ndTriangle(xL, yL, fringeSize, distance)

				// in order to have a perfect hexagon,
				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y1 + fringeSize
			}
			xs1 := []int{x12, x11, x12}
			ys1 := []int{y11, y12, y13}
			fill2 := draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+1+lines)%15))
			canvas.Polygon(xs1, ys1, fill2)

			xs1 = mirrorCoordinates(xs1, lines, fringeSize, 0)
			canvas.Polygon(xs1, ys1, fill2)
		}
	}
	canvas.End()
}
