package isogrids

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Isogrids is the handler for /isogrids/:key
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Isogrids(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	if key, err = route.Context.Get(r, "key"); err != nil {
		log.Println("Unable to get 'key' value: ", err)
		key = ""
	}

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	theme := extract.Theme(r)
	colorMap := colors.MapOfColorThemes()

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

	size := extract.Size(r)
	write.ImageSVG(w)
	draw.Isogrids(w, key, bg, fg, size)
}

func GridBW(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorPatterns()
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
	draw.IsogridsBW(w, "", bg, fg, size)
}

func Grid2Colors(w http.ResponseWriter, r *http.Request) {

	colorMap := colors.MapOfColorPatterns()
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
	draw.Isogrids2Colors(w, "", bg, fg, size)
}
