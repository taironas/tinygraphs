package squares

import (
	"image"
	"image/color"
)

// Palette builds an image with all the colors present in the theme array.
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
