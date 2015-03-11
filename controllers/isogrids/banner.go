package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// BannerRandom handler for /isogrids/banner/random.
// Generates a random banner isogrid image.
func BannerRandom(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)

	colorMap := colors.MapOfColorThemes()
	bg, fg := extract.ExtraColors(r, colorMap)
	theme := extract.Theme(r)
	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	}

	var colors []color.RGBA
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			numColors := extract.NumColors(r)
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}

	if newColors, err := extract.Colors(r); err == nil {
		colors = newColors
	}

	prob := extract.Probability(r, 1/float64(len(colors)))

	xt := extract.XTriangles(r)
	write.ImageSVG(w)
	isogrids.Random(w, "", colors, width, height, xt, prob)
}

// BannerRandomGradinet handler for /isogrids/banner/random/gradient.
// Generates a random gradient banner isogrid image.
func BannerRandomGradient(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)

	colorMap := colors.MapOfColorThemes()
	bg, fg := extract.ExtraColors(r, colorMap)
	theme := extract.Theme(r)
	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	}

	var colors []color.RGBA
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			numColors := extract.NumColors(r)
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}

	if newColors, err := extract.Colors(r); err == nil {
		colors = newColors
	}

	xt := extract.XTriangles(r)
	write.ImageSVG(w)
	isogrids.RandomGradient(w, "", colors, width, height, xt)
}
