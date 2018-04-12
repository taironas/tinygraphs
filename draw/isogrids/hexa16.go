package isogrids

import (
	"errors"
	"image/color"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// Hexa16 builds an image with lines x lines grids of half diagonals in the form of an hexagon
// it draws 6 triangles, triangle 1 to 5 are all rotations of triangle 0.
// triangle zero triangle on the center left.
func Hexa16(w io.Writer, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines
	distance := distanceTo3rdPoint(fringeSize)
	lines = size / fringeSize
	offset := ((fringeSize - distance) * lines) / 2

	fillTriangle := triangleColors(0, key, colors, lines)
	transparent := fillTransparent()

	isLeft := func(v int) bool { return (v % 2) == 0 }
	isRight := func(v int) bool { return (v % 2) != 0 }

	for xL := 0; xL < lines/2; xL++ {
		for yL := 0; yL < lines; yL++ {

			if isOutsideHexagon(xL, yL, lines) {
				continue
			}

			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}

			xs := []int{x2 + offset, x1 + offset, x2 + offset}
			ys := []int{y1, y2, y3}

			if fill, err := canFill(xL, yL, fillTriangle, isLeft, isRight); err != nil {
				canvas.Polygon(xs, ys, transparent)
			} else {
				canvas.Polygon(xs, ys, fill)
			}

			xsMirror := mirrorCoordinates(xs, lines, distance, offset*2)
			xLMirror := lines - xL - 1
			yLMirror := yL

			if fill, err := canFill(xLMirror, yLMirror, fillTriangle, isLeft, isRight); err != nil {
				canvas.Polygon(xsMirror, ys, transparent)
			} else {
				canvas.Polygon(xsMirror, ys, fill)
			}

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

			xs1 := []int{x12 + offset, x11 + offset, x12 + offset}
			ys1 := []int{y11, y12, y13}

			// triangles that go to the right

			if fill, err := canFill(xL, yL, fillTriangle, isRight, isLeft); err != nil {
				canvas.Polygon(xs1, ys1, transparent)
			} else {
				canvas.Polygon(xs1, ys1, fill)
			}

			xs1 = mirrorCoordinates(xs1, lines, distance, offset*2)

			if fill, err := canFill(xLMirror, yLMirror, fillTriangle, isRight, isLeft); err != nil {
				canvas.Polygon(xs1, ys1, transparent)
			} else {
				canvas.Polygon(xs1, ys1, fill)
			}
		}
	}
	canvas.End()
}

// triangleColors returns an array of strings, one for each sub triangle.
// Each string corresponds to an svg color.
// Colors are selected from the array of colors passed as parameter and the key.
func triangleColors(id int, key string, colors []color.RGBA, lines int) (tColors []string) {

	for _, t := range triangles[id] {
		x := t.x
		y := t.y
		tColors = append(tColors, draw.FillFromRGBA(draw.PickColor(key, colors, (x+3*y+lines)%15)))
	}
	return
}

// canFill returns a fill svg string given position. the fill is computed to be a rotation of the
// triangle 0 with the 'fills' array given as param.
func canFill(x, y int, fills []string, isLeft func(x int) bool, isRight func(x int) bool) (string, error) {
	l := newTrianglePosition(x, y, left)
	r := newTrianglePosition(x, y, right)

	if isLeft(x) && l.isInTriangle() {
		rid := l.rotationID()
		return fills[rid], nil
	} else if isRight(x) && r.isInTriangle() {
		rid := r.rotationID()
		return fills[rid], nil
	}
	return "", errors.New("cannot find svg fill for given position")
}
