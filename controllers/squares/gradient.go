package squares

import (
	"net/http"

	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// BannerGradient handler for "labs/squares/banner/gradient"
// generates a color gradient random grid image.
func BannerGradient(w http.ResponseWriter, r *http.Request) {

	width := extract.Width(r)
	height := extract.Height(r)
	xsquares := extract.XSquares(r)
	gv := extract.GradientVector(r, uint8(0), uint8(0), uint8(width), uint8(0))

	gColors := extract.GColors(r)
	colors := extract.Colors(r)
	prob := extract.Probability(r, 1/float64(len(gColors)))

	write.ImageSVG(w)
	squares.RandomGradientColorSVG(w, colors, gColors, gv, width, height, xsquares, prob)
}
