package isogrids

import (
	"image/color"
	"net/http"

	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func Skeleton(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
	var err error
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap[0][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap[0][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsSkeleton(w, "", bg, fg, size)
}
