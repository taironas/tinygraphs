package squares

import (
	"crypto/md5"
	"fmt"
	"github.com/ajstarks/svgo"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Square is the handler for /squares/[A-Za-z0-9]+/?
// build a 6x6 grid with alternate colors based on the number passed in the url
func Square(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 2); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {

		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

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
			canvas := svg.New(w)
			draw.SquaresSVG(canvas, key, bg, fg, size)
			write.ImageSVG(w, canvas)
		}
	}
}

// Color is the handler for /square/[0-8]/[a-zA-Z0-9]+/?
// build a 6x6 grid with alternate colors based on the number passed in the url
func Color(w http.ResponseWriter, r *http.Request) {

	if colorId, err := misc.PermalinkID(r, 2); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		if id, err1 := misc.PermalinkString(r, 3); err1 == nil {

			h := md5.New()
			io.WriteString(h, id)
			key := fmt.Sprintf("%x", h.Sum(nil)[:])
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
				canvas := svg.New(w)
				draw.SquaresSVG(canvas, key, colorMap[int(colorId)][0], colorMap[int(colorId)][1], size)
				write.ImageSVG(w, canvas)
			}
		} else {
			log.Printf("error when extracting permalink string: %v", err)
		}
	}
}
