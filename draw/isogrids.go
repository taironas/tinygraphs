package draw

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
)

// IsogridsRandom creates an isogrids image with half diagonals.
func IsogridsRandom(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
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
			canvas.Polygon(xs, ys, fillFromRGBA(randomColorFromArray(colors)))

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
			canvas.Polygon(xs1, ys1, fillFromRGBA(randomColorFromArray(colors)))
		}
	}
	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func IsogridsRandomMirror(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
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
			fill1 := fillFromRGBA(randomColorFromArray(colors))
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
			fill2 := fillFromRGBA(randomColorFromArray(colors))
			canvas.Polygon(xs1, ys1, fill2)
			// apply mirror:
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

// Isogrids builds an image with 10x10 grids of half diagonals
func Isogrids(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
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
			fill1 := fillFromRGBA(colorFromKeyAndArray(key, colors, (xL+3*yL+lines)%15))
			canvas.Polygon(xs, ys, fill1)
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
			fill2 := fillFromRGBA(colorFromKeyAndArray(key, colors, (xL+3*yL+1+lines)%15))
			canvas.Polygon(xs1, ys1, fill2)
			// apply mirror:
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

// IsogridsHexa builds an image with 10x10 grids of half diagonals
func IsogridsHexa(w http.ResponseWriter, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines

	// triangle grid here:
	for xL := -1; xL < lines/2; xL++ {
		for yL := -1; yL <= lines; yL++ {

			fill1 := fillWhite()
			fill2 := fillWhite()
			if isFill1InHexagon(xL, yL, lines) {
				fill1 = fillFromRGBA(colorFromKeyAndArray(key, colors, (xL+3*yL+lines)%15))
			}
			if isFill2InHexagon(xL, yL, lines) {
				fill2 = fillFromRGBA(colorFromKeyAndArray(key, colors, (xL+3*yL+1+lines)%15))
			}

			if !isFill1InHexagon(xL, yL, lines) && !isFill2InHexagon(xL, yL, lines) {
				continue
			}

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

			if lines%4 != 0 {
				xs[0] = x2
				xs[1] = x1
				xs[2] = x2
			}

			canvas.Polygon(xs, ys, fill1)

			// apply mirror:
			xs[0] = (lines * fringeSize) - xs[0]
			xs[1] = (lines * fringeSize) - xs[1]
			xs[2] = (lines * fringeSize) - xs[2]
			canvas.Polygon(xs, ys, fill1)

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
			if lines%4 != 0 {
				xs1[0] = x12
				xs1[1] = x11
				xs1[2] = x12
			}

			canvas.Polygon(xs1, ys1, fill2)
			xs1[0] = (lines * fringeSize) - xs1[0]
			xs1[1] = (lines * fringeSize) - xs1[1]
			xs1[2] = (lines * fringeSize) - xs1[2]
			canvas.Polygon(xs1, ys1, fill2)
		}
	}
	canvas.End()
}

func isFill1InHexagon(xL, yL, lines int) bool {
	if lines%6 == 0 {
		half := lines / 2
		start := half / 2
		if xL < start+1 {
			if yL > start-1 && yL < start+half+1 {
				return true
			}
		}
		if xL == half-1 {
			if yL > start-1-1 && yL < start+half+1+1 {
				return true
			}
		}
		return false
	} else if lines%4 == 0 {
		if xL == 0 {
			if yL > 1 && yL < 6 {
				return true
			}
		}
		if xL == 1 || xL == 2 {
			if yL > 0 && yL < 7 {
				return true
			}
		}
		if xL == 3 {
			if yL >= 0 && yL <= 7 {
				return true
			}
		}
		return false
	}
	return false
}

func isFill2InHexagon(xL, yL, lines int) bool {
	if lines%6 == 0 {
		half := lines / 2
		start := half / 2

		if xL < start {
			if yL > start-1 && yL < start+half {
				return true
			}
		}
		if xL == 1 {
			if yL > start-1-1 && yL < start+half+1 {
				return true
			}
		}
		if xL == half-1 {
			if yL > start-1-1 && yL < start+half+1 {
				return true
			}
		}
	} else if lines%4 == 0 {
		if xL == 0 || xL == 1 {
			if yL > 0 && yL < 6 {
				return true
			}
		}
		if xL == 1 {
			if yL > 0 && yL < 6 {
				return true
			}
		}
		if xL == 2 || xL == 3 {
			if yL >= 0 && yL <= 6 {
				return true
			}
		}
	}
	return false
}

func fillWhite() string {
	return "fill:rgb(255,255,255)"

}
