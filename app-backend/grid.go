package main

import (
	"image"
	"image/color"
	"log"
	"net/http"
)

// gridColorHandler is the handler for /grid/[0-8]
// build a 6x6 grid with alternate colors based on the number passed in the url
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

// grid6X6Handler is the handler for /grid/
// build a 6x6 grid with alternate black and white colors.
func grid6X6Handler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	color1 := color.RGBA{uint8(255), uint8(255), 255, 255}
	color2 := color.RGBA{uint8(0), uint8(0), 0, 255}
	drawGrid6X6(m, color1, color2)
	var img image.Image = m
	writeImage(w, &img)
}

// gradientHandler is the handler for /gradient/
func gradientHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	drawGradient(m)
	var img image.Image = m
	writeImage(w, &img)
}
