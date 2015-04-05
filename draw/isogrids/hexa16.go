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

// Hexa builds an image with lines x lines grids of half diagonals in the form of an hexagon
func Hexa16(w http.ResponseWriter, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines
	distance := distanceTo3rdPoint(fringeSize)
	lines = size / fringeSize
	offset := ((fringeSize - distance) * lines) / 2

	t1 := [][]int{
		{0, 1, right},
		{0, 2, right},
		{0, 3, right},
		{0, 2, left},
		{0, 3, left},
		{1, 2, right},
		{1, 3, right},
		{1, 2, left},
		{2, 2, right},
	}

	fillTriangle := []string{}
	for _, t := range t1 {
		x := t[0]
		y := t[1]
		fillTriangle = append(fillTriangle, draw.FillFromRGBA(draw.PickColor(key, colors, (x+3*y+lines)%15)))
	}

	for xL := 0; xL < lines/2; xL++ {
		for yL := 0; yL < lines; yL++ {

			fill1 := fillTransparent()
			fill2 := fillTransparent()

			if !isFill1InHexagon(xL, yL, lines) && !isFill2InHexagon(xL, yL, lines) {
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

			if (xL%2) == 0 && isInTriangleL(triangleId(xL, yL, left), xL, yL) {
				tid := triangleId(xL, yL, left)
				stid := subTriangleId(xL, yL, left, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[8])
				} else {
					canvas.Polygon(xs, ys, "fill:rgb(255,255,0)")
				}
			} else if (xL%2) != 0 && isInTriangleR(triangleId(xL, yL, right), xL, yL) {
				tid := triangleId(xL, yL, right)
				stid := subTriangleId(xL, yL, right, tid)

				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xs, ys, fillTriangle[8])
				} else {
					canvas.Polygon(xs, ys, "fill:rgb(255,255,0)")
				}
			} else {
				canvas.Polygon(xs, ys, fill1)
			}

			xsMirror := mirrorCoordinates(xs, lines, distance, offset*2)
			xLMirror := lines - xL - 1
			yLMirror := yL
			if (xLMirror%2) == 0 && isInTriangleL(triangleId(xLMirror, yLMirror, left), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, left)
				stid := subTriangleId(xLMirror, yLMirror, left, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[8])
				} else {
					canvas.Polygon(xsMirror, ys, "fill:rgb(255,255,0)")
				}
			} else if (xLMirror%2) != 0 && isInTriangleR(triangleId(xLMirror, yLMirror, right), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, right)
				stid := subTriangleId(xLMirror, yLMirror, right, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[8])

				} else {
					canvas.Polygon(xsMirror, ys, "fill:rgb(255,255,0)")
				}
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
			if (xL%2) != 0 && isInTriangleL(triangleId(xL, yL, left), xL, yL) {
				tid := triangleId(xL, yL, left)
				stid := subTriangleId(xL, yL, left, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[8])
				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}
			} else if (xL%2) == 0 && isInTriangleR(triangleId(xL, yL, right), xL, yL) {
				tid := triangleId(xL, yL, right)
				stid := subTriangleId(xL, yL, right, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[8])
				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}
			} else {
				canvas.Polygon(xs1, ys1, fill2)
			}

			xs1 = mirrorCoordinates(xs1, lines, distance, offset*2)
			if (xL%2) == 0 && isInTriangleL(triangleId(xLMirror, yLMirror, left), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, left)
				stid := subTriangleId(xLMirror, yLMirror, left, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[8])

				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}
			} else if (xL%2) != 0 && isInTriangleR(triangleId(xLMirror, yLMirror, right), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, right)
				stid := subTriangleId(xLMirror, yLMirror, right, tid)
				if SubTriangleIdsFromId(0)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else if SubTriangleIdsFromId(1)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[1])
				} else if SubTriangleIdsFromId(2)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[2])
				} else if SubTriangleIdsFromId(3)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[3])
				} else if SubTriangleIdsFromId(4)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[4])
				} else if SubTriangleIdsFromId(5)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[5])
				} else if SubTriangleIdsFromId(6)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[6])
				} else if SubTriangleIdsFromId(7)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[7])
				} else if SubTriangleIdsFromId(8)[tid] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[8])

				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}
			} else {
				canvas.Polygon(xs1, ys1, fill2)
			}
		}
	}
	canvas.End()
}

func isInTriangleL(id, xL, yL int) bool {
	if id == 0 {
		if (yL == 2 && xL == 0) ||
			(yL == 3 && xL == 0) {
			return true
		}
		if xL == 1 && yL == 2 {
			return true
		}
	} else if id == 1 {
		if yL == 1 && xL == 0 {
			return true
		}
		if (yL == 0 && xL == 1) ||
			(yL == 1 && xL == 1) {
			return true
		}
		if (yL == 0 && xL == 2) ||
			(yL == 1 && xL == 2) ||
			(yL == 2 && xL == 2) {
			return true
		}
	} else if id == 2 {
		if (xL == 3 && yL == 0) ||
			(xL == 3 && yL == 1) {
			return true
		}
		if xL == 4 && yL == 1 {
			return true
		}
	} else if id == 3 {
		if yL == 2 && xL == 3 {
			return true
		}
		if (yL == 2 && xL == 4) ||
			(yL == 3 && xL == 4) {
			return true
		}
		if (yL == 1 && xL == 5) ||
			(yL == 2 && xL == 5) ||
			(yL == 3 && xL == 5) {
			return true
		}
	} else if id == 4 {
		if (xL == 3 && yL == 3) ||
			(xL == 3 && yL == 4) {
			return true
		}
		if xL == 4 && yL == 4 {
			return true
		}
	} else if id == 5 {
		if yL == 4 && xL == 0 {
			return true
		}
		if (yL == 3 && xL == 1) ||
			(yL == 4 && xL == 1) {
			return true
		}
		if (yL == 3 && xL == 2) ||
			(yL == 4 && xL == 2) ||
			(yL == 5 && xL == 2) {
			return true
		}
	}
	return false
}

func isInTriangleR(id, xL, yL int) bool {
	if id == 0 {
		if (yL == 1 && xL == 0) ||
			(yL == 2 && xL == 0) ||
			(yL == 3 && xL == 0) {
			return true
		}
		if (yL == 2 && xL == 1) ||
			(yL == 3 && xL == 1) {
			return true
		}
		if yL == 2 && xL == 2 {
			return true
		}
	} else if id == 1 {
		if yL == 1 && xL == 1 {
			return true
		} else if (yL == 0 && xL == 2) ||
			(yL == 1 && xL == 2) {
			return true
		}
	} else if id == 2 {
		if (yL == 0 && xL == 3) ||
			(yL == 1 && xL == 3) ||
			(yL == 2 && xL == 3) {
			return true
		}
		if (yL == 0 && xL == 4) ||
			(yL == 1 && xL == 4) {
			return true
		}
		if yL == 1 && xL == 5 {
			return true
		}
	} else if id == 3 {
		if yL == 2 && xL == 4 {
			return true
		} else if (yL == 2 && xL == 5) ||
			(yL == 3 && xL == 5) {
			return true
		}
	} else if id == 4 {
		if (yL == 3 && xL == 3) ||
			(yL == 4 && xL == 3) ||
			(yL == 5 && xL == 3) {
			return true
		}
		if (yL == 3 && xL == 4) ||
			(yL == 4 && xL == 4) {
			return true
		}
		if yL == 4 && xL == 5 {
			return true
		}
	} else if id == 5 {
		if yL == 4 && xL == 1 {
			return true
		} else if (yL == 3 && xL == 2) ||
			(yL == 4 && xL == 2) {
			return true
		}
	}

	return false
}

func triangleId(x, y, direction int) int {

	triangles := [][]trianglePosition{
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

	triangles := [][]trianglePosition{
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

	for _, t := range triangles {
		for i, ti := range t {
			if ti.x == x && ti.y == y && (direction == ti.direction) {
				return i
			}
		}
	}

	return -1
}

func SubTriangleIdsFromId(lookforSubTriangleId int) []int {

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
