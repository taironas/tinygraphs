package draw

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
)

// Isogrids builds an image with 10x10 grids of half diagonals
func IsogridsRandom(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// triangle grid here:
	for xL := -1; xL <= lines; xL++ {
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
			canvas.Polygon(xs, ys, fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(randomColor(color1, color2))))

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
			canvas.Polygon(xs1, ys1, fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(randomColor(color1, color2))))
		}
	}
	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func IsogridsRandomMirror(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// triangle grid here:
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
			fill1 := fillFromRGBA(randomColor(color1, color2))
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
			fill2 := fillFromRGBA(randomColor(color1, color2))
			canvas.Polygon(xs1, ys1, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill2))
			// apply mirror:
			xs[0] = (lines * fringeSize) - xs[0]
			xs[1] = (lines * fringeSize) - xs[1]
			xs[2] = (lines * fringeSize) - xs[2]
			xs1[0] = (lines * fringeSize) - xs1[0]
			xs1[1] = (lines * fringeSize) - xs1[1]
			xs1[2] = (lines * fringeSize) - xs1[2]
			canvas.Polygon(xs, ys, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill1))
			canvas.Polygon(xs1, ys1, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill2))
		}
	}
	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func Isogrids(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// triangle grid here:
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
			fill1 := fillFromRGBA(colorFromKey(key, color1, color2, (xL+3*yL+lines)%15))
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
			fill2 := fillFromRGBA(colorFromKey(key, color1, color2, (xL+3*yL+1+lines)%15))
			canvas.Polygon(xs1, ys1, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill2))
			// apply mirror:
			xs[0] = (lines * fringeSize) - xs[0]
			xs[1] = (lines * fringeSize) - xs[1]
			xs[2] = (lines * fringeSize) - xs[2]
			xs1[0] = (lines * fringeSize) - xs1[0]
			xs1[1] = (lines * fringeSize) - xs1[1]
			xs1[2] = (lines * fringeSize) - xs1[2]
			canvas.Polygon(xs, ys, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill1))
			canvas.Polygon(xs1, ys1, fmt.Sprintf("stroke:black;stroke-width:2; %s", fill2))
		}
	}
	canvas.End()
}
