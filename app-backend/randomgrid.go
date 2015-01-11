package main

import (
	"image"
	"log"
	"net/http"
)

// handler for "/gird/random"
// generates a black and white grid random image.
func gridRandomHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	colorMap := MapOfColorPatterns()
	drawRandomGrid6X6(m, colorMap[0][0], colorMap[0][1])
	var img image.Image = m
	writeImage(w, &img)
}

// handler for "/grid/random/[0-9]"
// generates a grid random image with a specific color based on the colorMap
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
