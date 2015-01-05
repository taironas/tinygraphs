package main

import (
	"image"
	"image/color"

	//	"image/draw"
	"net/http"
)

func grid6X6Handler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	drawGrid6X6(m)
	var img image.Image = m
	writeImage(w, &img)
}

func gradientHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	drawGradient(m)
	var img image.Image = m
	writeImage(w, &img)
}

func drawGrid6X6(m *image.RGBA) {
	size := m.Bounds().Size()
	quad := size.X / 6
	for x := 0; x < size.X; x++ {
		val := ((x / quad) % quad) % 2
		for y := 0; y < size.Y; y++ {
			val2 := ((y / quad) % quad) % 2
			if val+val2 == 1 {
				color := color.RGBA{
					uint8(255),
					uint8(255),
					255,
					255}
				m.Set(x, y, color)
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
