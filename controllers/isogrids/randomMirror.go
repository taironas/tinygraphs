package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func RandomMirror(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorThemes()
	size := extract.Size(r)
	numColors := extract.NumColors(r)
	bg, fg := extract.ExtraColors(r, colorMap)

	theme := extract.Theme(r)
	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	}

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

	write.ImageSVG(w)
	isogrids.RandomMirror(w, "", colors, size)
}
