package isogrids

import (
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

const (
	left = iota
	right
)

var (
	triangles = [][]trianglePosition{
		[]trianglePosition{
			{0, 1, right},
			{0, 2, right},
			{0, 3, right},
			{0, 2, left},
			{0, 3, left},
			{1, 2, right},
			{1, 3, right},
			{1, 2, left},
			{2, 2, right},
		},
		[]trianglePosition{
			{0, 1, left},
			{1, 1, right},
			{1, 0, left},
			{1, 1, left},
			{2, 0, right},
			{2, 1, right},
			{2, 0, left},
			{2, 1, left},
			{2, 2, left},
		}, []trianglePosition{
			{3, 0, right},
			{3, 1, right},
			{3, 2, right},
			{3, 0, left},
			{3, 1, left},
			{4, 0, right},
			{4, 1, right},
			{4, 1, left},
			{5, 1, right},
		},
		[]trianglePosition{
			{3, 2, left},
			{4, 2, right},
			{4, 2, left},
			{4, 3, left},
			{5, 2, right},
			{5, 3, right},
			{5, 1, left},
			{5, 2, left},
			{5, 3, left},
		},
		[]trianglePosition{
			{3, 3, right},
			{3, 4, right},
			{3, 5, right},
			{3, 3, left},
			{3, 4, left},
			{4, 3, right},
			{4, 4, right},
			{4, 4, left},
			{5, 4, right},
		},
		[]trianglePosition{
			{0, 4, left},
			{1, 4, right},
			{1, 3, left},
			{1, 4, left},
			{2, 3, right},
			{2, 4, right},
			{2, 3, left},
			{2, 4, left},
			{2, 5, left},
		},
	}
)

// Hexa builds an image with lines x lines grids of half diagonals in the form of an hexagon
func Hexa16(w http.ResponseWriter, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines
	distance := distanceTo3rdPoint(fringeSize)
	lines = size / fringeSize
	offset := ((fringeSize - distance) * lines) / 2

	fillTriangle := triangleColors(0, key, colors, lines)

	for xL := 0; xL < lines/2; xL++ {
		for yL := 0; yL < lines; yL++ {

			if !isFill1InHexagon(xL, yL, lines) && !isFill2InHexagon(xL, yL, lines) {
				continue
			}

			fill1 := fillTransparent()
			fill2 := fillTransparent()

			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}

			xs := []int{x2 + offset, x1 + offset, x2 + offset}
			ys := []int{y1, y2, y3}

			if (xL%2) == 0 && isInTriangle(triangleId(xL, yL, left), xL, yL, left) {
				rid := rotationId(xL, yL, left)
				canvas.Polygon(xs, ys, fillTriangle[rid])
			} else if (xL%2) != 0 && isInTriangle(triangleId(xL, yL, right), xL, yL, right) {
				rid := rotationId(xL, yL, right)
				canvas.Polygon(xs, ys, fillTriangle[rid])
			} else {
				canvas.Polygon(xs, ys, fill1)
			}

			xsMirror := mirrorCoordinates(xs, lines, distance, offset*2)
			xLMirror := lines - xL - 1
			yLMirror := yL
			if (xLMirror%2) == 0 && isInTriangle(triangleId(xLMirror, yLMirror, left), xLMirror, yLMirror, left) {
				rid := rotationId(xLMirror, yLMirror, left)
				canvas.Polygon(xsMirror, ys, fillTriangle[rid])
			} else if (xLMirror%2) != 0 && isInTriangle(triangleId(xLMirror, yLMirror, right), xLMirror, yLMirror, right) {
				rid := rotationId(xLMirror, yLMirror, right)
				canvas.Polygon(xsMirror, ys, fillTriangle[rid])
			} else {
				canvas.Polygon(xsMirror, ys, fill1)
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
			if (xL%2) != 0 && isInTriangle(triangleId(xL, yL, left), xL, yL, left) {
				rid := rotationId(xL, yL, left)
				canvas.Polygon(xs1, ys1, fillTriangle[rid])
			} else if (xL%2) == 0 && isInTriangle(triangleId(xL, yL, right), xL, yL, right) {
				rid := rotationId(xL, yL, right)
				canvas.Polygon(xs1, ys1, fillTriangle[rid])
			} else {
				canvas.Polygon(xs1, ys1, fill2)
			}

			xs1 = mirrorCoordinates(xs1, lines, distance, offset*2)
			if (xL%2) == 0 && isInTriangle(triangleId(xLMirror, yLMirror, left), xLMirror, yLMirror, left) {
				rid := rotationId(xLMirror, yLMirror, left)
				canvas.Polygon(xs1, ys1, fillTriangle[rid])
			} else if (xL%2) != 0 && isInTriangle(triangleId(xLMirror, yLMirror, right), xLMirror, yLMirror, right) {
				rid := rotationId(xLMirror, yLMirror, right)
				canvas.Polygon(xs1, ys1, fillTriangle[rid])
			} else {
				canvas.Polygon(xs1, ys1, fill2)
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

// isInTriangle tells you whether the position x, y
// is in the triangle with id: 'id' if the sub triangle is a left or right one
// depending on the direction passed as param.
func isInTriangle(id, xL, yL, direction int) bool {
	if id == -1 {
		return false
	}
	for _, t := range triangles[id] {
		if t.direction != direction {
			continue
		}
		if t.x == xL && t.y == yL {
			return true
		}
	}
	return false
}

func triangleId(x, y, direction int) int {

	for i, t := range triangles {
		for _, ti := range t {
			if ti.x == x && ti.y == y && (direction == ti.direction) {
				return i
			}
		}
	}
	return -1
}

type trianglePosition struct {
	x, y, direction int
}

func subTriangleId(x, y, direction, id int) int {

	for _, t := range triangles {
		for i, ti := range t {
			if ti.x == x && ti.y == y && (direction == ti.direction) {
				return i
			}
		}
	}

	return -1
}

func subTriangleRotations(lookforSubTriangleId int) []int {

	m := map[int][]int{
		0: []int{0, 6, 8, 8, 2, 0},
		1: []int{1, 2, 5, 7, 6, 3},
		2: []int{2, 0, 0, 6, 8, 8},
		3: []int{3, 4, 7, 5, 4, 1},
		4: []int{4, 1, 3, 4, 7, 5},
		5: []int{5, 7, 6, 3, 1, 2},
		6: []int{6, 3, 1, 2, 5, 7},
		7: []int{7, 5, 4, 1, 3, 4},
		8: []int{8, 8, 2, 0, 0, 6},
	}
	if v, ok := m[lookforSubTriangleId]; ok {
		return v
	}
	return nil
}

// rotationId returns the original sub triangle id
// if the current triangle was rotated to position 0.
func rotationId(xL, yL, direction int) int {
	current_tid := triangleId(xL, yL, direction)
	current_stid := subTriangleId(xL, yL, direction, current_tid)
	numberOfSubTriangles := 9
	for i := 0; i < numberOfSubTriangles; i++ {
		rotations := subTriangleRotations(i)
		if rotations[current_tid] == current_stid {
			return i
		}
	}
	return -1
}
