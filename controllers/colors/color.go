package colors

import (
	"github.com/taironas/tinygraphs/write"
	"image"
	"image/color"
	"image/draw"
	"net/http"
)

func Black(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	var img image.Image = m
	write.Image(w, &img)
}

func Green(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	green := color.RGBA{0, 128, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	var img image.Image = m
	write.Image(w, &img)
}

func Blue(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	write.Image(w, &img)
}

func Red(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{255, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	write.ImageWithTemplate(w, &img)
}
