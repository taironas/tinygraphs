package squares

import (
	"image"
	"image/color"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

//Grid builds an image with 6X6 quadrants of alternate colors.
func Grid(m *image.RGBA, color1, color2 color.RGBA) {
	size := m.Bounds().Size()
	quad := size.X / 6
	for x := 0; x < size.X; x++ {
		val := (x / quad) % 2
		for y := 0; y < size.Y; y++ {
			val2 := (y / quad) % 2
			q := (val + val2) % 2
			if q == 0 {
				m.Set(x, y, color1)
			} else {
				m.Set(x, y, color2)
			}
		}
	}
}

// GridSVG builds an image with 6 by 6 quadrants of alternate colors.
func GridSVG(w http.ResponseWriter, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)
	squares := 6
	quadrantSize := size / squares
	colorMap := make(map[int]color.RGBA)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)

		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			if _, ok := colorMap[xQ]; !ok {
				if (xQ+yQ)%2 == 0 {
					colorMap[xQ] = color1
				} else {
					colorMap[xQ] = color2
				}
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, draw.FillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}

// Squares builds an image with 6 by 6 quadrants of alternate colors.
func Squares(m *image.RGBA, key string, colors []color.RGBA) {
	size := m.Bounds().Size()
	squares := 6
	quad := size.X / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	var currentYQuadrand = 0
	for y := 0; y < size.Y; y++ {
		yQuadrant := y / quad
		if yQuadrant != currentYQuadrand {
			// when y quadrant changes, clear map
			colorMap = make(map[int]color.RGBA)
			currentYQuadrand = yQuadrant
		}
		for x := 0; x < size.X; x++ {
			xQuadrant := x / quad
			if _, ok := colorMap[xQuadrant]; !ok {
				if float64(xQuadrant) < middle {
					colorMap[xQuadrant] = draw.PickColor(key, colors, xQuadrant+3*yQuadrant)
				} else if xQuadrant < squares {
					colorMap[xQuadrant] = colorMap[squares-xQuadrant-1]
				} else {
					colorMap[xQuadrant] = colorMap[0]
				}
			}
			m.Set(x, y, colorMap[xQuadrant])
		}
	}
}

// SquaresSVG builds an image with 6 by 6 quadrants of alternate colors.
func SquaresSVG(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	squares := 6
	quadrantSize := size / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)

		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			if _, ok := colorMap[xQ]; !ok {
				if float64(xQ) < middle {
					colorMap[xQ] = draw.PickColor(key, colors, xQ+3*yQ)
				} else if xQ < squares {
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorMap[xQ] = colorMap[0]
				}
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, draw.FillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}
