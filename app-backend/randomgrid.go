package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"net/http"
)

func gridRandomHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	colorMap := MapOfColorPatterns()
	drawRandomGrid6X6(m, colorMap[0][0], colorMap[0][1])
	var img image.Image = m
	writeImage(w, &img)
}

func gridRandomColorHandler(w http.ResponseWriter, r *http.Request) {
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

func drawRandomGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
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
				colorMap[yQuadrant] = getRandomColor(color1, color2)
			}
			m.Set(x, y, colorMap[yQuadrant])
		}
	}
}

func getRandomColor(c1, c2 color.RGBA) color.RGBA {
	r := rand.Intn(2)
	if r == 1 {
		return c1
	}
	return c2
}
