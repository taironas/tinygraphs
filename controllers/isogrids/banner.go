package isogrids

import (
	"image/color"
	"net/http"

	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// BannerRandom handler for /isogrids/banner/random.
// Generates a random banner isogrid image.
func BannerRandom(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)

	colors := GetColors(r)
	prob := extract.Probability(r, 1/float64(len(colors)))

	xt := extract.XTriangles(r)
	write.ImageSVG(w)
	isogrids.Random(w, "", colors, width, height, xt, prob)
}

func GetColors(r *http.Request) []color.RGBA {

	if newColors, err := extract.Colors(r); err == nil {
		return newColors
	}

	var colors []color.RGBA
	th := extract.Theme(r)

	if th == "base" {
		bg, fg := extract.ExtraColors(r)
		colors = append(colors, bg, fg)
	} else {
		m := tgColors.MapOfColorThemes()
		if _, ok := m[th]; ok {
			n := extract.NumColors(r)
			colors = append(colors, m[th][0:n]...)
		} else {
			colors = append(colors, m["base"]...)
		}
	}

	return colors
}

// BannerRandomGradinet handler for /isogrids/banner/random/gradient.
// Generates a random gradient banner isogrid image.
func BannerRandomGradient(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)

	colorMap := tgColors.MapOfColorThemes()
	bg, fg := extract.ExtraColors(r)
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
