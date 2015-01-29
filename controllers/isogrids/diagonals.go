package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Diagonals is the handler for /isogrids/labs/diagonals
// builds a full diagonal grid.
func Diagonals(w http.ResponseWriter, r *http.Request) {

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
	isogrids.Diagonals(w, "", bg, fg, size)
}

// Diagonals is the handler for /isogrids/labs/diagonals
// builds a half diagonal (each diagonal goes to the middle of the square) grid.
func HalfDiagonals(w http.ResponseWriter, r *http.Request) {

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
	isogrids.HalfDiagonals(w, "", bg, fg, size)
}
