package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Diagonals is the handler for /isogrids/labs/diagonals
// builds a full diagonal grid.
func Diagonals(w http.ResponseWriter, r *http.Request) {

	bg, fg := extract.ExtraColors(r)
	size := extract.Size(r)
	write.ImageSVG(w)
	isogrids.Diagonals(w, "", bg, fg, size)
}

// Diagonals is the handler for /isogrids/labs/diagonals
// builds a half diagonal (each diagonal goes to the middle of the square) grid.
func HalfDiagonals(w http.ResponseWriter, r *http.Request) {

	bg, fg := extract.ExtraColors(r)
	size := extract.Size(r)
	write.ImageSVG(w)
	isogrids.HalfDiagonals(w, "", bg, fg, size)
}
