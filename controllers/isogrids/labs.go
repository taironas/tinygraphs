package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// GridBW is the handler for /isogrids/labs/gridbw
// builds a 10x10 grid that alternate black and white colors.
func GridBW(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorThemes()
	var err error
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap["base"][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap["base"][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsBW(w, "", bg, fg, size)
}

// Grid2Colors is the handler for /isogrids/labs/grid2colors
// builds a 10x10 grid that alternate black and white colors.
func Grid2Colors(w http.ResponseWriter, r *http.Request) {

	size := extract.Size(r)
	write.ImageSVG(w)
	draw.Isogrids2Colors(w, "", size)
}
