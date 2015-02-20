package squares

import (
	"image"
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// RandomGrid builds a grid image with with x colors selected at random for each quadrant.
func RandomGrid(m *image.RGBA, colors []color.RGBA, xSquares int) {
	size := m.Bounds().Size()
	quad := size.X / xSquares
	colorMap := make(map[int]color.RGBA)
	var currentQuadrand = 0
	for x := 0; x < size.X; x++ {
		if x/quad != currentQuadrand {

			colorMap = make(map[int]color.RGBA)
			currentQuadrand = x / quad
		}
		for y := 0; y < size.Y; y++ {
			yQuadrant := y / quad
			if _, ok := colorMap[yQuadrant]; !ok {
				colorMap[yQuadrant] = draw.RandomColorFromArray(colors)
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}

// RandomGridSVG builds a grid image with with x colors selected at random for each quadrant.
func RandomGridSVG(w http.ResponseWriter, colors []color.RGBA, width, height, xSquares int) {
	canvas := svg.New(w)
	canvas.Start(width, height)
	squares := xSquares
	quadrantSize := width / squares
	colorMap := make(map[int]color.RGBA)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)

		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			if _, ok := colorMap[xQ]; !ok {
				colorMap[xQ] = draw.RandomColorFromArray(colors)
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, draw.FillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}

// RandomGradientGrid builds a grid image with with x colors selected at random for each quadrant going from brighter to dracker color.
func RandomGradientGrid(m *image.RGBA, colors []color.RGBA, xSquares int) {
	size := m.Bounds().Size()
	quad := size.X / xSquares
	colorMap := make(map[int]color.RGBA)
	var currentQuadrand = 0
	for x := 0; x < size.X; x++ {
		if x/quad != currentQuadrand {

			colorMap = make(map[int]color.RGBA)
			currentQuadrand = x / quad
		}
		percentage := 100 - int(float64(x)/float64(size.X)*100)
		for y := 0; y < size.Y; y++ {
			yQuadrant := y / quad
			if _, ok := colorMap[yQuadrant]; !ok {
				colorMap[yQuadrant] = draw.ColorByPercentage(colors, percentage)
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}

// RandomGradientGridSVG builds a grid image with with x colors selected at random for each quadrant.
func RandomGradientGridSVG(w http.ResponseWriter, colors []color.RGBA, width, height, xSquares int) {
	canvas := svg.New(w)
	canvas.Start(width, height)
	squares := xSquares
	quadrantSize := width / squares
	colorMap := make(map[int]color.RGBA)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)

		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			if _, ok := colorMap[xQ]; !ok {
				percentage := 100 - int(float64(xQ)/float64(squares)*100)
				colorMap[xQ] = draw.ColorByPercentage(colors, percentage)
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, draw.FillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}
