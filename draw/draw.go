package draw

import (
	"image"
	"image/color"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
)

//Grid6X6 builds an image with 6X6 quadrants of alternate colors.
func Grid6X6(m *image.RGBA, color1, color2 color.RGBA) {
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

// Grid6X6SVG builds an image with 6X6 quadrants of alternate colors.
func Grid6X6SVG(w http.ResponseWriter, color1, color2 color.RGBA, size int) {
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
			canvas.Rect(x, y, quadrantSize, quadrantSize, fillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}

// RandomGrid6X6 builds a grid image with with 2 colors selected at random for each quadrant.
func RandomGrid6X6(m *image.RGBA, colors []color.RGBA) {
	size := m.Bounds().Size()
	quad := size.X / 6
	colorMap := make(map[int]color.RGBA)
	var currentQuadrand = 0
	for x := 0; x < size.X; x++ {
		if x/quad != currentQuadrand {
			// quadrant changed, clear map
			colorMap = make(map[int]color.RGBA)
			currentQuadrand = x / quad
		}
		for y := 0; y < size.Y; y++ {
			yQuadrant := y / quad
			if _, ok := colorMap[yQuadrant]; !ok {
				colorMap[yQuadrant] = randomColorFromArray(colors)
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}

// RandomGrid6X6SVG builds a grid image with with 2 colors selected at random for each quadrant.
func RandomGrid6X6SVG(w http.ResponseWriter, colors []color.RGBA, size int) {
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
				colorMap[xQ] = randomColorFromArray(colors)
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, fillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}

// RandomSymetricInYGrid6X6 builds a grid image with with 2 colors selected at random for each quadrant.
func RandomSymetricInYGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
	size := m.Bounds().Size()
	squares := 6
	quad := size.X / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	var currentQuadrand = 0
	for x := 0; x < size.X; x++ {
		if x/quad != currentQuadrand {
			// when x quadrant changes, clear map
			colorMap = make(map[int]color.RGBA)
			currentQuadrand = x / quad
		}
		for y := 0; y < size.Y; y++ {
			yQuadrant := y / quad
			if _, ok := colorMap[yQuadrant]; !ok {
				if float64(yQuadrant) < middle {
					colorMap[yQuadrant] = randomColor(color1, color2)
				} else {
					colorMap[yQuadrant] = colorMap[squares-yQuadrant-1]
				}
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}

// RandomSymetricInXGrid6X6 builds a grid image with with 2 colors selected at random for each quadrant.
func RandomSymetricInXGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
	size := m.Bounds().Size()
	squares := 6
	quad := size.X / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	var currentQuadrand = 0
	for y := 0; y < size.Y; y++ {
		if y/quad != currentQuadrand {
			// when y quadrant changes, clear map
			colorMap = make(map[int]color.RGBA)
			currentQuadrand = y / quad
		}
		for x := 0; x < size.X; x++ {
			xQuadrant := x / quad
			if _, ok := colorMap[xQuadrant]; !ok {
				if float64(xQuadrant) < middle {
					colorMap[xQuadrant] = randomColor(color1, color2)
				} else {
					colorMap[xQuadrant] = colorMap[squares-xQuadrant-1]
				}
			}
			m.Set(x, y, colorMap[xQuadrant])
		}
	}
}

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
					colorMap[xQuadrant] = colorFromKeyAndArray(key, colors, xQuadrant+3*yQuadrant)
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
					colorMap[xQ] = colorFromKeyAndArray(key, colors, xQ+3*yQ)
				} else if xQ < squares {
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorMap[xQ] = colorMap[0]
				}
			}
			canvas.Rect(x, y, quadrantSize, quadrantSize, fillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}
