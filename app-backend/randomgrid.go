package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"net/http"
)

func randomColorGridHandler(w http.ResponseWriter, r *http.Request) {
	intID, err := PermalinkID(r, 2)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		var c1, c2 color.RGBA
		if intID == 1 {
			c1 = color.RGBA{uint8(0), uint8(144), 167, 255}
			c2 = color.RGBA{uint8(0), uint8(214), 244, 255}
		}
		if intID == 2 {
			c1 = color.RGBA{uint8(232), uint8(70), 134, 255}
			c2 = color.RGBA{uint8(181), uint8(181), 181, 255}
		}
		if intID == 3 {
			c1 = color.RGBA{uint8(255), uint8(237), 81, 255}
			c2 = color.RGBA{uint8(255), uint8(253), 136, 255}
		}
		if intID == 4 {
			c1 = color.RGBA{uint8(177), uint8(192), 24, 255}
			c2 = color.RGBA{uint8(86), uint8(165), 150, 255}
		}
		if intID == 5 {
			c1 = color.RGBA{uint8(208), uint8(2), 120, 255}
			c2 = color.RGBA{uint8(255), uint8(0), 118, 255}
		}
		if intID == 6 {
			c1 = color.RGBA{uint8(0), uint8(44), 47, 255}
			c2 = color.RGBA{uint8(126), uint8(176), 119, 255}
		}
		if intID == 7 {
			c1 = color.RGBA{uint8(29), uint8(24), 18, 255}
			c2 = color.RGBA{uint8(234), uint8(225), 219, 255}
		}
		if intID == 8 {
			c1 = color.RGBA{uint8(33), uint8(30), 26, 255}
			c2 = color.RGBA{uint8(176), uint8(209), 194, 255}
		}

		drawRandomGrid6X6(m, c1, c2)
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
