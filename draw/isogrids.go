package draw

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"image/color"
	"log"
	"net/http"
)

// IsogridsSVG builds an image with 10x10 grids of alternate colors.
func IsogridsSVG(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	size = 440
	canvas.Start(size, size)

	lines := 11
	fringeSize := size / lines
	// middle := math.Ceil(float64(lines) / float64(2))
	// colorMap := make(map[int]color.RGBA)

	// log.Printf(style)

	// horizontal lines
	for xL := 0; xL < lines; xL++ {
		x := xL * fringeSize
		lastY := (lines - 1) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) == 0 {

			//lastY = int(math.Ceil(float64(lastY) - float64(fringeSize)/float64(2)))
			lastY = lastY - fringeSize/2
			canvas.Line(x, 0, x, lastY, style)
		} else {
			// yNext := int(math.Ceil(float64(fringeSize) / float64(2)))
			yNext := fringeSize / 2
			canvas.Line(x, yNext, x, lastY, style)
		}
	}

	// diagonal lines, x --> y left down
	for xL := 0; xL < lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:orange;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) == 0 {
			xPrev := 0
			yPrev := (xL / 2) * fringeSize
			log.Println(yPrev)
			if yPrev > 0 {
				canvas.Line(xPrev, yPrev, x, 0, style)
			}
		}
	}

	// diagonal lines, x --> y right down
	for xL := 0; xL < lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:fuchsia;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) == 0 {
			xNext := (lines - 1) * fringeSize
			yNext := ((lines / 2) - (xL / 2)) * fringeSize
			if yNext > 0 {
				canvas.Line(x, 0, xNext, yNext, style)
			}
		}
	}

	// diagonal lines, x --> y left up
	for xL := 0; xL < lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:red;stroke-width:2; %s", fillFromRGBA(color2))
		if (xL % 2) != 0 {
			xLast := x
			yLast := (lines - 1) * fringeSize
			yPrev := ((lines - 1 - xL) + xL/2) * fringeSize
			if yPrev < (lines)*fringeSize {
				canvas.Line(xLast, yLast, 0, yPrev, style)
			}
		}
	}

	// diagonal lines, y --> x down
	xLast := (lines - 1) * fringeSize
	for yL := 0; yL < lines; yL++ {
		style := fmt.Sprintf("stroke:red;stroke-width:2; %s", fillFromRGBA(color2))
		y := yL * fringeSize
		yDiag := y + (5 * fringeSize)
		if yDiag < ((lines - 1) * fringeSize) {
			canvas.Line(0, y, xLast, yDiag, style)
		}
	}

	// diagonal lines, y --> x up
	for yL := 0; yL < lines; yL++ {
		style := fmt.Sprintf("stroke:blue;stroke-width:2; %s", fillFromRGBA(color2))
		y := yL * fringeSize
		yDiag := y - (5 * fringeSize)
		if yDiag > 0 {
			canvas.Line(0, y, xLast, yDiag, style)
		}
	}

	// diagonal lines, x --> y left
	for yL := 0; yL < lines; yL++ {
		style := fmt.Sprintf("stroke:green;stroke-width:2; %s", fillFromRGBA(color2))
		y := yL * fringeSize
		yDiag := y - (lines / 2 * fringeSize)
		if yDiag > 0 {
			canvas.Line(0, y, xLast, yDiag, style)
		}
	}

	// for yL := 0; yL < lines; yL++ {
	// 	y := yL * fringeSize
	// 	colorMap = make(map[int]color.RGBA)
	// 	// <line x1="10" x2="50" y1="110" y2="150" stroke="orange" fill="transparent" stroke-width="5"/>
	// 	for xL := 0; xL < lines; xL++ {
	// 		x := xL * fringeSize
	// 		// lastX := (lines-1)*fringeSize
	// 		lastY := (lines - 1) * fringeSize

	// 		canvas.Line(x, y, x, lastY)
	// 		// 	if _, ok := colorMap[xQ]; !ok {
	// 		// 		if float64(xQ) < middle {
	// 		// 			colorMap[xQ] = colorFromKey(key, color1, color2, xQ+3*yQ)
	// 		// 		} else if xQ < squares {
	// 		// 			colorMap[xQ] = colorMap[squares-xQ-1]
	// 		// 		} else {
	// 		// 			colorMap[xQ] = colorMap[0]
	// 		// 		}
	// 		// 	}
	// 		// 	canvas.Rect(x, y, quadrantSize, quadrantSize, fillFromRGBA(colorMap[xQ]))
	// 	}
	// }
	canvas.End()
}
