package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// GridBW is the handler for /isogrids/labs/gridbw
// builds a 10x10 grid that alternate black and white colors.
func GridBW(w http.ResponseWriter, r *http.Request) {

	bg, fg := extract.ExtraColors(r, colors.MapOfColorThemes())
	size := extract.Size(r)
	write.ImageSVG(w)
	isogrids.BlackWhite(w, "", bg, fg, size)
}

// Grid2Colors is the handler for /isogrids/labs/grid2colors
// builds a 10x10 grid that alternate black and white colors.
func Grid2Colors(w http.ResponseWriter, r *http.Request) {

	size := extract.Size(r)
	write.ImageSVG(w)
	isogrids.TwoColors(w, "", size)
}
