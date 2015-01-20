package draw

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
)

// Diagonals builds an image with 10x10 grids of normal diagonals.
func Diagonals(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines OK
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		canvas.Line(x, 0, x, lastY, style)
	}

	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x, 0, style)
		}
	}

	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x, 0, style)
		}
	}

	canvas.End()
}

// HalfDiagonals builds an image with 10x10 grids of half diagonals
func HalfDiagonals(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		canvas.Line(x, 0, x, lastY, style)
	}

	// y -- > x up right
	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}
	// x --> y down right
	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func IsogridsSkeleton(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		firstY := 0
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) != 0 {
			lastY = lastY - fringeSize/2
			firstY = fringeSize / 2
		}
		canvas.Line(x, firstY, x, lastY, style)
	}

	// y -- > x up right
	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}
	// x --> y down right
	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func IsogridsBW(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		firstY := 0
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) != 0 {
			lastY = lastY - fringeSize/2
			firstY = fringeSize / 2
		}
		canvas.Line(x, firstY, x, lastY, style)
	}

	// y -- > x up right
	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}
	// x --> y down right
	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

	// triangle grid here:
	for xL := 0; xL <= lines; xL++ {
		for yL := 0; yL <= lines; yL++ {
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
			canvas.Polygon(xs, ys, fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2)))
		}
	}

	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func Isogrids2Colors(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		firstY := 0
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) != 0 {
			lastY = lastY - fringeSize/2
			firstY = fringeSize / 2
		}
		canvas.Line(x, firstY, x, lastY, style)
	}

	// y -- > x up right
	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}
	// x --> y down right
	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

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
			canvas.Polygon(xs, ys, fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2)))

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
			canvas.Polygon(xs1, ys1, fmt.Sprintf("stroke:black;stroke-width:2; fill: rgb(31, 71, 176);")) //, fillFromRGBA(color2)))

		}
	}

	canvas.End()
}

// Isogrids builds an image with 10x10 grids of half diagonals
func IsogridsRandom(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		firstY := 0
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) != 0 {
			lastY = lastY - fringeSize/2
			firstY = fringeSize / 2
		}
		canvas.Line(x, firstY, x, lastY, style)
	}

	// y -- > x up right
	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}
	// x --> y down right
	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

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
	size = 400
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	// vertical lines
	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		firstY := 0
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) != 0 {
			lastY = lastY - fringeSize/2
			firstY = fringeSize / 2
		}
		canvas.Line(x, firstY, x, lastY, style)
	}

	// y -- > x up right
	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}
	// x --> y down right
	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

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
