package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func GridBW(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorPatterns()
	var err error
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap[0][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap[0][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsBW(w, "", bg, fg, size)
}

func Grid2Colors(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorPatterns()
	var err error
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap[0][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap[0][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.Isogrids2Colors(w, "", bg, fg, size)
}
