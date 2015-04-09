package isogrids

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/cache"
	"github.com/taironas/tinygraphs/draw/isogrids"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// Hexa is the handler for /isogrids/hexa/:key
// builds an hexagon with alternate colors.
func Hexa(w http.ResponseWriter, r *http.Request) {
	var err error
	size := extract.Size(r)

	var key string
	if key, _ = route.Context.Get(r, "key"); err != nil {
		key = ""
	}
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	lines := extract.Hexalines(r)
	colors := extract.Colors(r)

	if Cache.IsCached(&w, r, key, colors, size) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	write.ImageSVG(w)
	isogrids.Hexa(w, key, colors, size, lines)
}
