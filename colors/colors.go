package colors

import (
	"image/color"
)

// MapOfColorPatterns is used to build random images with colors that combine together.
func MapOfColorPatterns() map[int][]color.RGBA {
	return map[int][]color.RGBA{
		0: []color.RGBA{
			color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
			color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
		},

		1: []color.RGBA{
			color.RGBA{uint8(0), uint8(144), uint8(167), uint8(255)},
			color.RGBA{uint8(0), uint8(214), uint8(244), uint8(255)},
		},
		2: []color.RGBA{
			color.RGBA{uint8(232), uint8(70), 134, 255},
			color.RGBA{uint8(181), uint8(181), 181, 255},
		},
		3: []color.RGBA{
			color.RGBA{uint8(255), uint8(237), 81, 255},
			color.RGBA{uint8(255), uint8(253), 136, 255},
		},
		4: []color.RGBA{
			color.RGBA{uint8(177), uint8(192), 24, 255},
			color.RGBA{uint8(86), uint8(165), 150, 255},
		},
		5: []color.RGBA{
			color.RGBA{uint8(208), uint8(2), 120, 255},
			color.RGBA{uint8(255), uint8(0), 118, 255},
		},
		6: []color.RGBA{
			color.RGBA{uint8(0), uint8(44), 47, 255},
			color.RGBA{uint8(126), uint8(176), 119, 255},
		},
		7: []color.RGBA{
			color.RGBA{uint8(29), uint8(24), 18, 255},
			color.RGBA{uint8(234), uint8(225), 219, 255},
		},
		8: []color.RGBA{
			color.RGBA{uint8(33), uint8(30), 26, 255},
			color.RGBA{uint8(176), uint8(209), 194, 255},
		},
	}
}
