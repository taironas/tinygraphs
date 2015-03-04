package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Gradient handler for "labs/isogrids/banner/gradient"
// generates a color gradient random grid image.
func BannerGradient(w http.ResponseWriter, r *http.Request) {

	width := extract.Width(r)
	height := extract.Height(r)
	xtriangles := extract.XTriangles(r)
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
	isogrids.RandomGradientSVG(w, colors, gColors, gv, width, height, xtriangles)
}
