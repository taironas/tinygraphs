package isogrids

import (
	"fmt"
	"image/color"
	"log"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

const (
	both = iota
	left
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
		{0, 2, left},
		{0, 3, right},
		{0, 3, left},
		{1, 2, right},
		{1, 2, left},
		{1, 3, right},
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

			if isFill1InHexagon(xL, yL, lines) {
				fill1 = "fill:rgb(0,0,0)" // draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+lines)%15))
			}
			if isFill2InHexagon(xL, yL, lines) {
				fill2 = "fill:rgb(0,0,0)" //draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+1+lines)%15))
			}

			if !isFill1InHexagon(xL, yL, lines) && !isFill2InHexagon(xL, yL, lines) {
				continue
			}
			if false {
				fmt.Printf(fill1, fill2)
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
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xs, ys, fillTriangle[0])
				} else {
					canvas.Polygon(xs, ys, "fill:rgb(255,255,0)")
				}
			} else if (xL%2) != 0 && isInTriangleR(triangleId(xL, yL, right), xL, yL) {
				tid := triangleId(xL, yL, right)
				stid := subTriangleId(xL, yL, right, tid)
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xs, ys, fillTriangle[0])
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
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[0])
				} else {
					canvas.Polygon(xsMirror, ys, "fill:rgb(255,255,0)")
				}
			} else if (xLMirror%2) != 0 && isInTriangleR(triangleId(xLMirror, yLMirror, right), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, right)
				stid := subTriangleId(xLMirror, yLMirror, right, tid)
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xsMirror, ys, fillTriangle[0])
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
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}
				// canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
			} else if (xL%2) == 0 && isInTriangleR(triangleId(xL, yL, right), xL, yL) {
				tid := triangleId(xL, yL, right)
				stid := subTriangleId(xL, yL, right, tid)
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}

				// canvas.Polygon(xs1, ys1, "fill:rgb(255,0,0)")
			} else {
				canvas.Polygon(xs1, ys1, fill2)
			}

			xs1 = mirrorCoordinates(xs1, lines, distance, offset*2)
			if (xL%2) == 0 && isInTriangleL(triangleId(xLMirror, yLMirror, left), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, left)
				stid := subTriangleId(xLMirror, yLMirror, left, tid)
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
				} else {
					canvas.Polygon(xs1, ys1, "fill:rgb(255,255,0)")
				}
			} else if (xL%2) != 0 && isInTriangleR(triangleId(xLMirror, yLMirror, right), xLMirror, yLMirror) {
				tid := triangleId(xLMirror, yLMirror, right)
				stid := subTriangleId(xLMirror, yLMirror, right, tid)
				stids := SubTriangleIdsFromId(9)
				if stids[tid-1] == stid {
					canvas.Polygon(xs1, ys1, fillTriangle[0])
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
	// id = 6
	// return false
	fmt.Printf("x: %d, y:%d, id:%d\n", xL, yL, id)
	if id == 1 {
		if (yL == 2 && xL == 0) ||
			(yL == 3 && xL == 0) {
			log.Println("yes")
			return true
		}
		if xL == 1 && yL == 2 {
			log.Println("yes")
			return true
		}
	} else if id == 2 {
		if yL == 1 && xL == 0 {
			log.Println("yes")
			return true
		}
		if (yL == 0 && xL == 1) ||
			(yL == 1 && xL == 1) {
			log.Println("yes")
			return true
		}
		if (yL == 0 && xL == 2) ||
			(yL == 1 && xL == 2) ||
			(yL == 2 && xL == 2) {
			log.Println("yes")
			return true
		}
	} else if id == 3 {
		if (xL == 3 && yL == 0) ||
			(xL == 3 && yL == 1) {
			log.Println("yes")
			return true
		}
		if xL == 4 && yL == 1 {
			log.Println("yes")
			return true
		}
	} else if id == 4 {
		if yL == 2 && xL == 3 {
			log.Println("yes")
			return true
		}
		if (yL == 2 && xL == 4) ||
			(yL == 3 && xL == 4) {
			log.Println("yes")
			return true
		}
		if (yL == 1 && xL == 5) ||
			(yL == 2 && xL == 5) ||
			(yL == 3 && xL == 5) {
			log.Println("yes")
			return true
		}
	} else if id == 5 {
		if (xL == 3 && yL == 3) ||
			(xL == 3 && yL == 4) {
			log.Println("yes")
			return true
		}
		if xL == 4 && yL == 4 {
			log.Println("yes")
			return true
		}
	} else if id == 6 {
		if yL == 4 && xL == 0 {
			log.Println("yes")
			return true
		}
		if (yL == 3 && xL == 1) ||
			(yL == 4 && xL == 1) {
			log.Println("yes")
			return true
		}
		if (yL == 3 && xL == 2) ||
			(yL == 4 && xL == 2) ||
			(yL == 5 && xL == 2) {
			log.Println("yes")
			return true
		}
	}
	return false
}

func isInTriangleR(id, xL, yL int) bool {
	// return false
	// id = 6
	fmt.Printf("x: %d, y:%d, id:%d\n", xL, yL, id)
	if id == 1 {
		if (yL == 1 && xL == 0) ||
			(yL == 2 && xL == 0) ||
			(yL == 3 && xL == 0) {
			log.Println("yes")
			return true
		}
		if (yL == 2 && xL == 1) ||
			(yL == 3 && xL == 1) {
			log.Println("yes")
			return true
		}
		if yL == 2 && xL == 2 {
			log.Println("yes")
			return true
		}
	} else if id == 2 {
		if yL == 1 && xL == 1 {
			log.Println("yes")
			return true
		} else if (yL == 0 && xL == 2) ||
			(yL == 1 && xL == 2) {
			return true
		}
	} else if id == 3 {
		if (yL == 0 && xL == 3) ||
			(yL == 1 && xL == 3) ||
			(yL == 2 && xL == 3) {
			log.Println("yes")
			return true
		}
		if (yL == 0 && xL == 4) ||
			(yL == 1 && xL == 4) {
			log.Println("yes")
			return true
		}
		if yL == 1 && xL == 5 {
			log.Println("yes")
			return true
		}
	} else if id == 4 {
		if yL == 2 && xL == 4 {
			log.Println("yes")
			return true
		} else if (yL == 2 && xL == 5) ||
			(yL == 3 && xL == 5) {
			return true
		}
	} else if id == 5 {
		if (yL == 3 && xL == 3) ||
			(yL == 4 && xL == 3) ||
			(yL == 5 && xL == 3) {
			log.Println("yes")
			return true
		}
		if (yL == 3 && xL == 4) ||
			(yL == 4 && xL == 4) {
			log.Println("yes")
			return true
		}
		if yL == 4 && xL == 5 {
			log.Println("yes")
			return true
		}
	} else if id == 6 {
		if yL == 4 && xL == 1 {
			log.Println("yes")
			return true
		} else if (yL == 3 && xL == 2) ||
			(yL == 4 && xL == 2) {
			return true
		}
	}

	return false
}

func triangleId(x, y, direction int) int {

	t1 := [][]int{
		{0, 1, right},
		{0, 2, both},
		{0, 3, both},
		{1, 2, both},
		{1, 3, right},
		{2, 2, right},
	}
	t2 := [][]int{
		{0, 1, left},
		{1, 0, left},
		{1, 1, both},
		{2, 0, both},
		{2, 1, both},
		{2, 2, left},
	}
	t3 := [][]int{
		{5, 1, right},
		{4, 0, right},
		{4, 1, both},
		{3, 0, both},
		{3, 1, both},
		{3, 2, right},
	}
	t4 := [][]int{
		{5, 1, left},
		{5, 2, both},
		{5, 3, both},
		{4, 2, both},
		{4, 3, left},
		{3, 2, left},
	}
	t5 := [][]int{
		{5, 4, right},
		{4, 3, right},
		{4, 4, both},
		{3, 3, both},
		{3, 4, both},
		{3, 5, right},
	}
	t6 := [][]int{
		{1, 4, both},
		{2, 3, both},
		{2, 4, both},
		{0, 4, left},
		{1, 3, left},
		{2, 5, left},
	}

	for _, p := range t1 {
		if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
			return 1
		}
	}
	for _, p := range t2 {
		if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
			return 2
		}
	}
	for _, p := range t3 {
		if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
			return 3
		}
	}
	for _, p := range t4 {
		if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
			return 4
		}
	}
	for _, p := range t5 {
		if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
			return 5
		}
	}

	for _, p := range t6 {
		if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
			return 6
		}
	}

	return -1
}

func subTriangleId(x, y, direction, id int) int {

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
	t2 := [][]int{
		{0, 1, left},
		{1, 1, right},
		{1, 0, left},
		{1, 1, left},
		{2, 0, right},
		{2, 1, right},
		{2, 0, left},
		{2, 1, left},
		{2, 2, left},
	}
	t3 := [][]int{
		{3, 0, right},
		{3, 1, right},
		{3, 2, right},
		{3, 0, left},
		{3, 1, left},
		{4, 0, right},
		{4, 1, right},
		{4, 1, left},
		{5, 1, right},
	}
	t4 := [][]int{
		{3, 2, left},
		{4, 2, right},
		{4, 2, left},
		{4, 3, left},
		{5, 2, right},
		{5, 3, right},
		{5, 1, left},
		{5, 2, left},
		{5, 3, left},
	}
	t5 := [][]int{
		{3, 3, right},
		{3, 4, right},
		{3, 5, right},
		{3, 3, left},
		{3, 4, left},
		{4, 3, right},
		{4, 4, right},
		{4, 4, left},
		{5, 4, right},
	}
	t6 := [][]int{
		{0, 4, left},
		{1, 4, right},
		{1, 3, left},
		{1, 4, left},
		{2, 3, right},
		{2, 4, right},
		{2, 3, left},
		{2, 4, left},
		{2, 5, left},
	}

	if id == 1 {
		for i, p := range t1 {
			if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
				return i + 1
			}
		}
	}

	if id == 2 {
		for i, p := range t2 {
			if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
				return i + 1
			}
		}
	}

	if id == 3 {
		for i, p := range t3 {
			if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
				return i + 1
			}
		}
	}

	if id == 4 {
		for i, p := range t4 {
			if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
				return i + 1
			}
		}
	}

	if id == 5 {
		for i, p := range t5 {
			if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
				return i + 1
			}
		}
	}

	if id == 6 {
		for i, p := range t6 {
			if p[0] == x && p[1] == y && (direction == p[2] || p[2] == both) {
				return i + 1
			}
		}
	}

	return -1
}

func SubTriangleIdsFromId(lookforSubTriangleId int) []int {

	m := map[int][]int{
		1: []int{1, 7, 9, 9, 3, 1},
		2: []int{2, 3, 6, 8, 7, 4},
		3: []int{3, 1, 1, 7, 9, 9},
		4: []int{4, 5, 8, 6, 5, 2},
		5: []int{5, 2, 4, 5, 8, 6},
		6: []int{6, 8, 7, 4, 2, 3},
		7: []int{7, 4, 2, 3, 6, 8},
		8: []int{8, 6, 5, 2, 4, 5},
		9: []int{9, 9, 3, 1, 1, 7},
	}
	if v, ok := m[lookforSubTriangleId]; ok {
		return v
	}
	return nil
}
