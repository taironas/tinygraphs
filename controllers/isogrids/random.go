package isogrids

import (
	"image/color"
	"log"
	"net/http"
	"strconv"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func Random(w http.ResponseWriter, r *http.Request) {

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
	draw.IsogridsRandom(w, "", bg, fg, size)
}

func RandomColor(w http.ResponseWriter, r *http.Request) {
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

	colorMap := colors.MapOfColorPatterns()
	bg, err1 := extract.Background(r)
	if err1 != nil {
		bg = colorMap[int(colorId)][0]
	}
	fg, err2 := extract.Foreground(r)
	if err2 != nil {
		fg = colorMap[int(colorId)][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsRandom(w, "", bg, fg, size)
}

func RandomMirror(w http.ResponseWriter, r *http.Request) {
	colorMap := colors.MapOfColorPatterns()
	bg, err1 := extract.Background(r)
	if err1 != nil {
		bg = colorMap[0][0]
	}
	fg, err2 := extract.Foreground(r)
	if err2 != nil {
		fg = colorMap[0][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsRandomMirror(w, "", bg, fg, size)
}

func RandomMirrorColor(w http.ResponseWriter, r *http.Request) {
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
	colorMap := colors.MapOfColorPatterns()
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap[int(colorId)][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap[int(colorId)][1]
	}
	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsRandomMirror(w, "", bg, fg, size)
}
