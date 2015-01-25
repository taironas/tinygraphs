package isogrids

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/taironas/route"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

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

	colorMap := tgColors.MapOfColorPatterns()
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap[0][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap[0][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.Isogrids(w, key, bg, fg, size)
}

// Color is the handler for /isogrids/:colorId/:key
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Color(w http.ResponseWriter, r *http.Request) {
	var err error
	var id string
	if id, err = route.Context.Get(r, "colorId"); err != nil {
		log.Println("Unable to get 'colorId' value: ", err)
		id = "0"
	}

	var colorId int64
	if colorId, err = strconv.ParseInt(id, 0, 64); err != nil {
		colorId = 0
	}

	var key string
	if key, err = route.Context.Get(r, "key"); err != nil {
		log.Println("Unable to get 'key' value: ", err)
		key = ""
	}

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	colorMap := tgColors.MapOfColorPatterns()
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap[int(colorId)][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap[int(colorId)][1]
	}

	size := extract.Size(r)
	write.ImageSVG(w)
	draw.Isogrids(w, key, bg, fg, size)
}

func GridBW(w http.ResponseWriter, r *http.Request) {

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
	draw.IsogridsBW(w, "", bg, fg, size)
}

func Grid2Colors(w http.ResponseWriter, r *http.Request) {

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
	draw.Isogrids2Colors(w, "", bg, fg, size)
}
