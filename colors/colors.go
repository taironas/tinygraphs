package colors

import (
	"image/color"
)

type GradientVector struct {
	X1, Y1, X2, Y2 uint8
}

func White() color.RGBA {
	return color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)}
}

func Black() color.RGBA {
	return color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
}

// MapOfColorThemes is used to build random images with colors that go together.
func MapOfColorThemes() map[string][]color.RGBA {
	return map[string][]color.RGBA{
		"base": {
			White(),
			Black(),
		},
		"sugarsweets": {
			{255, 245, 249, 255}, //background
			{232, 70, 134, 255},  // main
			{232, 70, 186, 255},  // 2dary
			{232, 70, 81, 255},   // 2dary
		},
		"heatwave": {
			{255, 251, 243, 255}, //background
			{251, 76, 1, 255},    // main
			{207, 36, 24, 255},   // 2dary
			{255, 141, 20, 255},  // 2dary
		},
		"daisygarden": {
			{254, 251, 230, 255}, //background
			{177, 192, 24, 255},  // main
			{135, 173, 22, 255},  // 2dary
			{235, 199, 23, 255},  // 2dary
		},
		"seascape": {
			{240, 245, 244, 255}, //background
			{0, 114, 122, 255},   // main
			{0, 81, 130, 255},    // 2dary
			{0, 171, 145, 255},   // 2dary
		},
		"summerwarmth": {
			{255, 255, 232, 255}, //background
			{247, 228, 64, 255},  // main
			{255, 207, 72, 255},  // 2dary
			{193, 222, 53, 255},  // 2dary
		},
		"bythepool": {
			{237, 242, 255, 255}, //background
			{80, 135, 243, 255},  // main
			{92, 186, 242, 255},  // 2dary
			{131, 103, 244, 255}, // 2dary
		},
		"duskfalling": {
			{244, 234, 252, 255}, //background
			{142, 74, 219, 255},  // main
			{70, 86, 212, 255},   // 2dary
			{201, 64, 206, 255},  // 2dary
		},
		"frogideas": {
			{226, 255, 222, 255}, //background
			{93, 214, 75, 255},   // main
			{67, 191, 134, 255},  // 2dary
			{148, 232, 56, 255},  // 2dary
		},
		"berrypie": {
			{255, 240, 240, 255}, //background
			{248, 0, 6, 255},     // main
			{245, 44, 118, 255},  // 2dary
			{255, 79, 0, 255},    // 2dary
		},
	}
}
