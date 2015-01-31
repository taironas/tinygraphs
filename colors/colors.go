package colors

import (
	"image/color"
)

// MapOfColorThemes is used to build random images with colors that go together.
func MapOfColorThemes() map[string][]color.RGBA {
	return map[string][]color.RGBA{
		"base": []color.RGBA{
			color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
			color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
		},

		"sugarsweets": []color.RGBA{
			color.RGBA{232, 70, 134, 255},  // main
			color.RGBA{255, 245, 249, 255}, //background
			color.RGBA{232, 70, 186, 255},  // 2dary
			color.RGBA{232, 70, 81, 255},   // 2dary
		},
		"heatwave": []color.RGBA{
			color.RGBA{251, 76, 1, 255},    // main
			color.RGBA{255, 251, 243, 255}, //background
			color.RGBA{207, 36, 24, 255},   // 2dary
			color.RGBA{255, 141, 20, 255},  // 2dary
		},
		"daisygarden": []color.RGBA{
			color.RGBA{177, 192, 24, 255},  // main
			color.RGBA{254, 251, 230, 255}, //background
			color.RGBA{135, 173, 22, 255},  // 2dary
			color.RGBA{235, 199, 23, 255},  // 2dary
		},
		"seascape": []color.RGBA{
			color.RGBA{0, 114, 122, 255},   // main
			color.RGBA{240, 245, 244, 255}, //background
			color.RGBA{0, 81, 130, 255},    // 2dary
			color.RGBA{0, 171, 145, 255},   // 2dary
		},
		"summerwarmth": []color.RGBA{
			color.RGBA{247, 228, 64, 255},  // main
			color.RGBA{255, 255, 232, 255}, //background
			color.RGBA{255, 207, 72, 255},  // 2dary
			color.RGBA{193, 222, 53, 255},  // 2dary
		},
		"bythepool": []color.RGBA{
			color.RGBA{80, 135, 243, 255},  // main
			color.RGBA{237, 242, 255, 255}, //background
			color.RGBA{92, 186, 242, 255},  // 2dary
			color.RGBA{131, 103, 244, 255}, // 2dary
		},
		"duskfalling": []color.RGBA{
			color.RGBA{142, 74, 219, 255},  // main
			color.RGBA{244, 234, 252, 255}, //background
			color.RGBA{70, 86, 212, 255},   // 2dary
			color.RGBA{201, 64, 206, 255},  // 2dary

		},
		"frogideas": []color.RGBA{
			color.RGBA{93, 214, 75, 255},   // main
			color.RGBA{226, 255, 222, 255}, //background
			color.RGBA{67, 191, 134, 255},  // 2dary
			color.RGBA{148, 232, 56, 255},  // 2dary

		},
		"berrypie": []color.RGBA{
			color.RGBA{248, 0, 6, 255},     // main
			color.RGBA{255, 240, 240, 255}, //background
			color.RGBA{245, 44, 118, 255},  // 2dary
			color.RGBA{255, 79, 0, 255},    // 2dary
		},
	}
}
