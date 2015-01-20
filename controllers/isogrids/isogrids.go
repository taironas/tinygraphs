package isogrids

import (
	"crypto/md5"
	"fmt"

	"io"
	"log"
	"net/http"

	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"

	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
)

// Isogrids is the handler for /isogrids/[a-zA-Z0-9]+/?.
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Isogrids(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 2); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

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
		draw.IsogridsSkeleton(w, key, bg, fg, size)
	}
}

func Skeleton(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

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
		draw.IsogridsSkeleton(w, key, bg, fg, size)
	}
}

func Diagonals(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

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
		draw.Diagonals(w, key, bg, fg, size)
	}
}

func HalfDiagonals(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

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
		draw.HalfDiagonals(w, key, bg, fg, size)
	}
}
