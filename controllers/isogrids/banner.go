package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// BannerRandom handler for /isogrids/banner/random.
// Generates a random banner isogrid image.
func BannerRandom(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)

	colors := extract.Colors(r)
	prob := extract.Probability(r, 1/float64(len(colors)))

	xt := extract.XTriangles(r)
	write.ImageSVG(w)
	isogrids.Random(w, colors, width, height, xt, prob)
}

// BannerRandomGradient handler for /isogrids/banner/random/gradient.
// Generates a random gradient banner isogrid image.
func BannerRandomGradient(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)

	colors := extract.Colors(r)

	xt := extract.XTriangles(r)
	write.ImageSVG(w)
	isogrids.RandomGradient(w, colors, width, height, xt)
}
