package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Random handler for /isogrids/random.
// Generates a random isogrid image.
func Random(w http.ResponseWriter, r *http.Request) {

	colors := extract.Colors(r)
	prob := extract.Probability(r, 1/float64(len(colors)))

	size := extract.Size(r)
	lines := extract.Lines(r)
	write.ImageSVG(w)
	isogrids.Random(w, colors, size, size, lines, prob)
}
