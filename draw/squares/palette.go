package squares

import (
	"image"
	"image/color"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// Palette builds an JPEG image with all the colors present in the theme color array.
func Palette(m *image.RGBA, theme []color.RGBA) {
	size := m.Bounds().Size()
	numColors := len(theme)
	quadrant := size.X / numColors
	for x := 0; x < size.X; x++ {
		currentQuadrant := (x / quadrant) % numColors
		for y := 0; y < size.Y; y++ {
			m.Set(x, y, theme[currentQuadrant])
		}
	}
}

// PaletteSVG builds an SVG image with all the colors present in the theme color array.
func PaletteSVG(w io.Writer, theme []color.RGBA, width, height int) {
	canvas := svg.New(w)
	canvas.Start(width, height)

	numColors := len(theme)
	quadrant := width / numColors

	for xQ := 0; xQ < numColors; xQ++ {
		x := xQ * quadrant
		currentQuadrant := (x / quadrant) % numColors
		canvas.Rect(x, 0, quadrant, height, draw.FillFromRGBA(theme[currentQuadrant]))
	}
	canvas.End()
}
