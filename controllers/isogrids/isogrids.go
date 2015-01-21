package isogrids

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"io"
	"net/http"

	"github.com/taironas/route"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"

	"github.com/taironas/tinygraphs/write"
)

// Isogrids is the handler for /isogrids/:key
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Isogrids(w http.ResponseWriter, r *http.Request) {
	key := route.Context.Get(r, "key")
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.Isogrids(w, key, bg, fg, size)
}

// Color is the handler for /isogrids/:colorId/:key
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Color(w http.ResponseWriter, r *http.Request) {
	id := route.Context.Get(r, "colorId")

	colorId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		colorId = 0
	}

	key := route.Context.Get(r, "key")
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.Isogrids(w, key, bg, fg, size)
}

func Skeleton(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.IsogridsSkeleton(w, "", bg, fg, size)
}

func Diagonals(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.Diagonals(w, "", bg, fg, size)
}

func HalfDiagonals(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.HalfDiagonals(w, "", bg, fg, size)
}

func GridBW(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.IsogridsBW(w, "", bg, fg, size)
}

func Grid2Colors(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.Isogrids2Colors(w, "", bg, fg, size)
}

func Random(w http.ResponseWriter, r *http.Request) {

	colorMap := tgColors.MapOfColorPatterns()
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
	draw.IsogridsRandom(w, "", bg, fg, size)
}

func RandomColor(w http.ResponseWriter, r *http.Request) {
	id := route.Context.Get(r, "colorId")
	colorId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		colorId = 0
	}
	colorMap := tgColors.MapOfColorPatterns()
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
	colorMap := tgColors.MapOfColorPatterns()
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
	id := route.Context.Get(r, "colorId")
	colorId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		colorId = 0
	}
	colorMap := tgColors.MapOfColorPatterns()
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
	draw.IsogridsRandomMirror(w, "", bg, fg, size)
}
