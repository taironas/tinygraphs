package isogrids

import (
	"crypto/md5"
	"fmt"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	// "github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"io"
	"log"
	"net/http"
	//	"strings"
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

		// e := `"` + key + `"`
		// w.Header().Set("Etag", e)
		// w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

		// if match := r.Header.Get("If-None-Match"); match != "" {
		// 	if strings.Contains(match, e) {
		// 		w.WriteHeader(http.StatusNotModified)
		// 		return
		// 	}
		// }

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
		// if f := extract.Format(r); f == format.SVG {
		write.ImageSVG(w)
		draw.IsogridsSVG(w, key, bg, fg, size)
		// }
	}
}
