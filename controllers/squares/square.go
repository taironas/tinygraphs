package squares

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// Square is the handler for /squares/:key
// builds a 6x6 grid with alternate colors based on the number passed in the url.
func Square(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	if key, err = route.Context.Get(r, "key"); err != nil {
		log.Println("Unable to get 'key' value: ", err)
		key = ""
	}
	theme := extract.Theme(r)

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	e := `"` + theme + key + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

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
