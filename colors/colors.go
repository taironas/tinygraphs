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

// MapOfColorThemes is used to build random images with colors that combine together.
func MapOfColorThemes() map[string][]color.RGBA {
	return map[string][]color.RGBA{
		"base": []color.RGBA{
			color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
			color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
		},

		"sugarsweets": []color.RGBA{
			color.RGBA{232, 70, 134, 255},  // main
			color.RGBA{232, 70, 186, 255},  // 2dary
			color.RGBA{232, 70, 81, 255},   // 2dary
			color.RGBA{255, 245, 249, 255}, //background
		},
		"heatwave": []color.RGBA{
			color.RGBA{251, 76, 1, 255},    // main
			color.RGBA{207, 36, 24, 255},   // 2dary
			color.RGBA{255, 141, 20, 255},  // 2dary
			color.RGBA{255, 251, 243, 255}, //background
		},
		"daisygarden": []color.RGBA{
			color.RGBA{177, 192, 24, 255},  // main
			color.RGBA{135, 173, 22, 255},  // 2dary
			color.RGBA{235, 199, 23, 255},  // 2dary
			color.RGBA{254, 251, 230, 255}, //background
		},
		"seascape": []color.RGBA{
			color.RGBA{0, 114, 122, 255},   // main
			color.RGBA{0, 81, 130, 255},    // 2dary
			color.RGBA{0, 171, 145, 255},   // 2dary
			color.RGBA{240, 245, 244, 255}, //background
		},
		"summerwarmth": []color.RGBA{
			color.RGBA{247, 228, 64, 255},  // main
			color.RGBA{255, 207, 72, 255},  // 2dary
			color.RGBA{193, 222, 53, 255},  // 2dary
			color.RGBA{255, 255, 232, 255}, //background
		},
		"bythepool": []color.RGBA{
			color.RGBA{80, 135, 243, 255},  // main
			color.RGBA{92, 186, 242, 255},  // 2dary
			color.RGBA{131, 103, 244, 255}, // 2dary
			color.RGBA{237, 242, 255, 255}, //background
		},
		"duskfalling": []color.RGBA{
			color.RGBA{142, 74, 219, 255},  // main
			color.RGBA{70, 86, 212, 255},   // 2dary
			color.RGBA{201, 64, 206, 255},  // 2dary
			color.RGBA{244, 234, 252, 255}, //background
		},
		"frogideas": []color.RGBA{
			color.RGBA{93, 214, 75, 255},   // main
			color.RGBA{67, 191, 134, 255},  // 2dary
			color.RGBA{148, 232, 56, 255},  // 2dary
			color.RGBA{226, 255, 222, 255}, //background
		},
		"berrypie": []color.RGBA{
			color.RGBA{248, 0, 6, 255},     // main
			color.RGBA{245, 44, 118, 255},  // 2dary
			color.RGBA{255, 79, 0, 255},    // 2dary
			color.RGBA{255, 240, 240, 255}, //background
		},
	}
}
