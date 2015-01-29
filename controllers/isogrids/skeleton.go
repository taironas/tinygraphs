package isogrids

import (
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func Skeleton(w http.ResponseWriter, r *http.Request) {

	fg, bg := extract.ExtraColors(r, colors.MapOfColorThemes())
	size := extract.Size(r)
	write.ImageSVG(w)
	isogrids.Skeleton(w, "", bg, fg, size)
}
