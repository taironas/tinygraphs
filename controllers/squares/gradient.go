package squares

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Gradient handler for "/labs/squares/gradient/:key"
// generates a color gradient random grid image.
func Gradient(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	if key, err = route.Context.Get(r, "key"); err != nil {
		log.Println("Unable to get 'key' value: ", err)
		key = ""
	}

	theme := extract.Theme(r)

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	numColors := extract.NumColors(r)
	colorMap := colors.MapOfColorThemes()

	bg, fg := extract.ExtraColors(r, colorMap)

	var colors []color.RGBA
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}

	size := extract.Size(r)
	write.ImageSVG(w)
	squares.GradientSVG(w, key, colors, size, size, 6)
}

// Gradient handler for "labs/squares/banner/gradient"
// generates a color gradient random grid image.
func BannerGradient(w http.ResponseWriter, r *http.Request) {

	width := extract.Width(r)
	height := extract.Height(r)
	xsquares := extract.XSquares(r)
	numColors := extract.NumColors(r)
	gv := extract.GradientVector(r, uint8(0), uint8(0), uint8(width), uint8(0))

	colorMap := colors.MapOfColorThemes()

	bg, fg := extract.ExtraColors(r, colorMap)

	var colors, gColors []color.RGBA
	theme := extract.Theme(r)
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			colors = append(colors, colorMap[theme][0:numColors]...)
			gColors = colorMap[theme][1:3]
		} else {
			colors = append(colors, colorMap["base"]...)
			gColors = colorMap[theme]
		}
	} else {
		colors = append(colors, bg, fg)
		gColors = []color.RGBA{bg, fg}
	}

	write.ImageSVG(w)
	squares.RandomGradientSVG(w, colors, gColors, gv, width, height, xsquares)
}
