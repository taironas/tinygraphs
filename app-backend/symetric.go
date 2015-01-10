package main

import (
	"image"
	"image/color"
	"log"
	"net/http"
)

// handler for "/gird/random/symetric"
// generates a black and white grid random image.
func gridRandomSymetricHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	colorMap := MapOfColorPatterns()
	drawRandomSymetricGrid6X6(m, colorMap[0][0], colorMap[0][1])
	var img image.Image = m
	writeImage(w, &img)
}

// handler for "/grid/random/symetric[0-9]"
// generates a grid random image with a specific color based on the colorMap
func gridRandomSymetricColorHandler(w http.ResponseWriter, r *http.Request) {
	intID, err := PermalinkID(r, 3)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		colorMap := MapOfColorPatterns()
		drawRandomGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		writeImage(w, &img)
	}
}

// drawRandomGrid6X6 builds a grid image with with 2 colors selected at random for each quadrant.
func drawRandomSymetricGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
	size := m.Bounds().Size()
	quad := size.X / 6
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
				colorMap[yQuadrant] = getRandomColor(color1, color2)
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}
