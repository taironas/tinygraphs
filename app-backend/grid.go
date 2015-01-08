package main

import (
	"image"
	"image/color"
	"log"
	"net/http"
)

func gridColorHandler(w http.ResponseWriter, r *http.Request) {
	intID, err := PermalinkID(r, 2)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		colorMap := MapOfColorPatterns()
		drawGrid6X6(m, colorMap[int(intID)][0], colorMap[int(intID)][1])
		var img image.Image = m
		writeImage(w, &img)
	}
}

func grid6X6Handler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	color1 := color.RGBA{uint8(255), uint8(255), 255, 255}
	color2 := color.RGBA{uint8(0), uint8(0), 0, 255}
	drawGrid6X6(m, color1, color2)
	var img image.Image = m
	writeImage(w, &img)
}

func gradientHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	drawGradient(m)
	var img image.Image = m
	writeImage(w, &img)
}

func drawGrid6X6(m *image.RGBA, color1, color2 color.RGBA) {
	size := m.Bounds().Size()
	quad := size.X / 6
	for x := 0; x < size.X; x++ {
		val := (x / quad) % 2
		for y := 0; y < size.Y; y++ {
			val2 := (y / quad) % 2
			q := (val + val2) % 2
			if q == 0 {
				m.Set(x, y, color1)
			} else {
				m.Set(x, y, color2)
			}
		}
	}
}

func drawGradient(m *image.RGBA) {
	size := m.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255}
			m.Set(x, y, color)
		}
	}
}
