package isogrids

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/draw/isogrids"
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

	colors := extract.Colors(r)
	size := extract.Size(r)
	lines := extract.Lines(r)

	write.ImageSVG(w)
	isogrids.Isogrids(w, key, colors, size, lines)
}
