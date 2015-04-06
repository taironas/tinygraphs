package isogrids

import (
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// Hexa builds an image with lines x lines grids of half diagonals in the form of an hexagon
func Hexa(w http.ResponseWriter, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines
	distance := distanceTo3rdPoint(fringeSize)
	lines = size / fringeSize
	offset := ((fringeSize - distance) * lines) / 2

	for xL := 0; xL < lines/2; xL++ {
		for yL := 0; yL < lines; yL++ {

			fill1 := fillTransparent()
			fill2 := fillTransparent()

			if isFill1InHexagon(xL, yL, lines) {
				fill1 = draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+lines)%15))
			}
			if isFill2InHexagon(xL, yL, lines) {
				fill2 = draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+1+lines)%15))
			}

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

			canvas.Polygon(xs, ys, fill1)

			xsMirror := mirrorCoordinates(xs, lines, distance, offset*2)
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

			xs1 := []int{x12 + offset, x11 + offset, x12 + offset}
			ys1 := []int{y11, y12, y13}

			canvas.Polygon(xs1, ys1, fill2)

			xs1 = mirrorCoordinates(xs1, lines, distance, offset*2)
			canvas.Polygon(xs1, ys1, fill2)
		}
	}
	canvas.End()
}

func isFill1InHexagon(xL, yL, lines int) bool {

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
}

func isFill2InHexagon(xL, yL, lines int) bool {

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
	return false
}

// isOutsideHexagon returns true if the current position is outside the hexagon.
func isOutsideHexagon(xL, yL, lines int) bool {
	return !isFill1InHexagon(xL, yL, lines) && !isFill2InHexagon(xL, yL, lines)
}

// fillWhite returns the svg style to paint an object white.
func fillWhite() string {
	return "fill:rgb(255,255,255)"

}

func fillTransparent() string {
	return "fill:rgb(255,255,255); fill-opacity:0"
}
