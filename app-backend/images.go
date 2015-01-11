package main

import (
	"image"
	"image/color"
	"image/draw"
	"net/http"
)

func blackHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func greenHandler(w http.ResponseWriter, r *http.Request) {
	// prepare
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	green := color.RGBA{0, 128, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{255, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImageWithTemplate(w, &img)
}
