package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"net/http"
)

// handler for "/gird/random/symetric/x"
// generates a black and white grid random image.
func gridRandomSymetricXHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	colorMap := MapOfColorPatterns()
	drawRandomSymetricInXGrid6X6(m, colorMap[0][0], colorMap[0][1])
	var img image.Image = m
	writeImage(w, &img)
}

// handler for "/gird/random/symetric/y"
// generates a black and white grid random image.
func gridRandomSymetricYHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	colorMap := MapOfColorPatterns()
	drawRandomSymetricInYGrid6X6(m, colorMap[0][0], colorMap[0][1])
	var img image.Image = m
	writeImage(w, &img)
}

// handler for "/grid/random/symetric/y/[0-9]"
// generates a grid random image with a specific color based on the colorMap
func gridRandomSymetricYColorHandler(w http.ResponseWriter, r *http.Request) {
	intID, err := PermalinkID(r, 5)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		colorMap := MapOfColorPatterns()
		drawRandomSymetricInYGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		writeImage(w, &img)
	}
}

// handler for "/grid/random/symetric/x/[0-9]"
// generates a grid random image with a specific color based on the colorMap
func gridRandomSymetricXColorHandler(w http.ResponseWriter, r *http.Request) {
	intID, err := PermalinkID(r, 5)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		colorMap := MapOfColorPatterns()
		drawRandomSymetricInXGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		writeImage(w, &img)
	}
}

// drawRandomGrid6X6 builds a grid image with with 2 colors selected at random for each quadrant.
func drawRandomSymetricInYGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
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
					colorMap[yQuadrant] = getRandomColor(color1, color2)
				} else {
					colorMap[yQuadrant] = colorMap[squares-yQuadrant-1] //getRandomColor(color1, color2)
				}
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}

// drawRandomGrid6X6 builds a grid image with with 2 colors selected at random for each quadrant.
func drawRandomSymetricInXGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
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
					colorMap[xQuadrant] = getRandomColor(color1, color2)
				} else {
					colorMap[xQuadrant] = colorMap[squares-xQuadrant-1]
				}
			}
			m.Set(x, y, colorMap[xQuadrant])
		}
	}
}
