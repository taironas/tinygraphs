package isogrids

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Hexa is the handler for /isogrids/hexa/:key
// builds an hexagon from a 10x10 grid with alternate colors.
func Hexa(w http.ResponseWriter, r *http.Request) {
	var err error
	colorMap := colors.MapOfColorThemes()
	size := extract.Size(r)

	var key string
	if key, _ = route.Context.Get(r, "key"); err != nil {
		key = ""
	}
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	theme := extract.Theme(r)
	numColors := extract.NumColors(r)

	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap["base"][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap["base"][1]
	}

	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	}

	var colors []color.RGBA
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}
	write.ImageSVG(w)
	draw.IsogridsHexa(w, key, colors, size)
}
