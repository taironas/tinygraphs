package isogrids

import (
	"image/color"
	"net/http"

	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func Skeleton(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorThemes()
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
	isogrids.Skeleton(w, "", bg, fg, size)
}
