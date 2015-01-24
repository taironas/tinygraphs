package squares

import (
	"crypto/md5"
	"fmt"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/taironas/route"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// Square is the handler for /squares/:key
// builds a 6x6 grid with alternate colors based on the number passed in the url.
func Square(w http.ResponseWriter, r *http.Request) {
	key, _ := route.Context.Get(r, "key")
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	e := `"` + key + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

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
	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.Squares(m, key, bg, fg)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		draw.SquaresSVG(w, key, bg, fg, size)
	}
}

// Color is the handler for /square/:colorId/:key
// build a 6x6 grid with alternate colors based on the number passed in the url
func Color(w http.ResponseWriter, r *http.Request) {

	id, _ := route.Context.Get(r, "colorId")
	colorId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		colorId = 0
	}

	key, _ := route.Context.Get(r, "key")
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])
	strId := strconv.FormatInt(colorId, 10)

	e := `"` + key + `-` + strId + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	size := extract.Size(r)
	colorMap := tgColors.MapOfColorPatterns()
	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.Squares(m, key, colorMap[int(colorId)][0], colorMap[int(colorId)][1])
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		draw.SquaresSVG(w, key, colorMap[int(colorId)][0], colorMap[int(colorId)][1], size)
	}
}
