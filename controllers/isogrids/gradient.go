package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Gradient handler for "labs/isogrids/banner/gradient"
// generates a color gradient random grid image.
func BannerGradient(w http.ResponseWriter, r *http.Request) {

	width := extract.Width(r)
	height := extract.Height(r)
	xtriangles := extract.XTriangles(r)
	gv := extract.GradientVector(r, uint8(0), uint8(0), uint8(width), uint8(0))

	colors := extract.Colors(r)
	gColors := extract.GColors(r)

	prob := extract.Probability(r, 1/float64(len(colors)))

	write.ImageSVG(w)
	isogrids.RandomGradientColor(w, colors, gColors, gv, width, height, xtriangles, prob)
}
