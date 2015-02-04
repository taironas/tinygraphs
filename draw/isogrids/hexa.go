package isogrids

import (
	"image/color"
	"log"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// Hexa builds an image with 10x10 grids of half diagonals
func Hexa(w http.ResponseWriter, key string, colors []color.RGBA, size, lines int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	fringeSize := size / lines
	log.Println(fringeSize)
	// distance of center of vector to third point of equilateral triangles
	// ABC triangle, O is the center of AB vector
	// OC = SQRT(AC^2 - AO^2)
	distance := int(math.Ceil(math.Sqrt((float64(fringeSize) * float64(fringeSize)) - (float64(fringeSize)/float64(2))*(float64(fringeSize)/float64(2)))))
	fringeSize = distance
	lines = size / fringeSize
	offset := size - fringeSize*lines
	log.Println(offset)
	for xL := 0; xL < lines/2; xL++ {
		for yL := 0; yL < lines; yL++ {

			fill1 := fillWhite()
			fill2 := fillWhite()
			if isFill1InHexagon(xL, yL, lines) {
				fill1 = draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+lines)%15))
			}
			if isFill2InHexagon(xL, yL, lines) {
				fill2 = draw.FillFromRGBA(draw.PickColor(key, colors, (xL+3*yL+1+lines)%15))
			}

			if !isFill1InHexagon(xL, yL, lines) && !isFill2InHexagon(xL, yL, lines) {
				continue
			}

			var x1, x2, x3, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1 = xL * fringeSize
				x2 = xL*fringeSize + distance
				x3 = x1
				y1 = yL * fringeSize
				y2 = y1 + fringeSize/2
				y3 = yL*fringeSize + distance
			} else {
				x1 = xL*fringeSize + distance
				x2 = xL * fringeSize
				x3 = x1
				y1 = yL * fringeSize
				y2 = y1 + fringeSize/2
				y3 = yL*fringeSize + distance
			}
			xs := []int{x1, x2, x3}
			ys := []int{y1, y2, y3}
			if lines%4 != 0 {
				xs[0] = x2
				xs[1] = x1
				xs[2] = x2
			}
			canvas.Polygon(xs, ys, fill1)

			xsMirror := []int{0, 0, 0}
			xsMirror[0] = (lines * fringeSize) - xs[0]
			xsMirror[1] = (lines * fringeSize) - xs[1]
			xsMirror[2] = (lines * fringeSize) - xs[2]

			canvas.Polygon(xsMirror, ys, fill1)

			var x11, x12, x13, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11 = xL*fringeSize + distance
				x12 = xL * fringeSize
				x13 = x11
				y11 = yL*fringeSize + fringeSize/2
				y12 = y3 // to have a perfect hexagon, we make sure that previews triangle and this one touch each other in this point
				y13 = yL*fringeSize + fringeSize/2 + distance
			} else {
				x11 = xL * fringeSize
				x12 = xL*fringeSize + distance
				x13 = x11
				y11 = yL*fringeSize + fringeSize/2
				y12 = y1 + fringeSize
				y13 = yL*fringeSize + fringeSize/2 + distance
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
