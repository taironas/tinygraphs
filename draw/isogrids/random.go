package isogrids

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

func Random(w http.ResponseWriter, key string, colors []color.RGBA, width, height, lines int) {
	canvas := svg.New(w)
	canvas.Start(width, height)

	fringeSize := width / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = width / fringeSize

	for xL := 0; xL < lines; xL++ {
		for yL := 0; yL < lines; yL++ {
			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}
			canvas.Polygon(xs, ys, draw.FillFromRGBA(draw.RandomColorFromArray(colors)))

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
			canvas.Polygon(xs1, ys1, draw.FillFromRGBA(draw.RandomColorFromArray(colors)))
		}
	}
	canvas.End()
}

// Random creates an isogrids svg image with half diagonals.
func RandomGradient(w http.ResponseWriter, key string, colors []color.RGBA, width, height, lines int) {
	canvas := svg.New(w)
	canvas.Start(width, height)

	fringeSize := width / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = width / fringeSize

	for xL := 0; xL < lines; xL++ {
		percentage := int(float64(xL) / float64(lines) * 100)
		for yL := 0; yL < lines; yL++ {
			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}
			canvas.Polygon(xs, ys, draw.FillFromRGBA(draw.ColorByPercentage(colors, percentage)))

			var x11, x12, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11, y11, x12, y12, _, y13 = left2ndTriangle(xL, yL, fringeSize, distance)

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
			canvas.Polygon(xs1, ys1, draw.FillFromRGBA(draw.ColorByPercentage(colors, percentage)))
		}
	}
	canvas.End()
}

// RandomMirror builds an image with 10x10 grids of half diagonals
func RandomMirror(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	for xL := -1; xL <= lines/2; xL++ {
		for yL := -1; yL <= lines; yL++ {
			var x1, x2, x3, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1 = (xL) * fringeSize
				x2 = (xL + 1) * fringeSize
				x3 = x1
				y1 = yL * fringeSize
				y2 = y1 + fringeSize/2
				y3 = (yL + 1) * fringeSize
			} else {
				x1 = (xL + 1) * fringeSize
				x2 = xL * fringeSize
				x3 = x1
				y1 = yL * fringeSize
				y2 = y1 + fringeSize/2
				y3 = (yL + 1) * fringeSize
			}
			xs := []int{x1, x2, x3}
			ys := []int{y1, y2, y3}
			fill1 := draw.FillFromRGBA(draw.RandomColorFromArray(colors))
			canvas.Polygon(xs, ys, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill1))
			var x11, x12, x13, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11 = (xL + 1) * fringeSize
				x12 = (xL) * fringeSize
				x13 = x11
				y11 = yL*fringeSize + fringeSize/2
				y12 = y11 + fringeSize/2
				y13 = (yL+1)*fringeSize + fringeSize/2
			} else {
				x11 = (xL) * fringeSize
				x12 = (xL + 1) * fringeSize
				x13 = x11
				y11 = yL*fringeSize + fringeSize/2
				y12 = y1 + fringeSize
				y13 = (yL+1)*fringeSize + fringeSize/2
			}
			xs1 := []int{x11, x12, x13}
			ys1 := []int{y11, y12, y13}
			fill2 := draw.FillFromRGBA(draw.RandomColorFromArray(colors))
			canvas.Polygon(xs1, ys1, fill2)

			xs[0] = (lines * fringeSize) - xs[0]
			xs[1] = (lines * fringeSize) - xs[1]
			xs[2] = (lines * fringeSize) - xs[2]
			xs1[0] = (lines * fringeSize) - xs1[0]
			xs1[1] = (lines * fringeSize) - xs1[1]
			xs1[2] = (lines * fringeSize) - xs1[2]
			canvas.Polygon(xs, ys, fill1)
			canvas.Polygon(xs1, ys1, fill2)
		}
	}
	canvas.End()
}
