package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func RandomMirror(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	colors := extract.Colors(r)
	prob := extract.Probability(r, 1/float64(len(colors)))

	write.ImageSVG(w)
	isogrids.RandomMirror(w, colors, size, prob)
}
