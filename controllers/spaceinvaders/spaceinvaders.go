package spaceinvaders

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/taironas/route"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/spaceinvaders"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// SpaceInvaders handler for /spaceinvaders/:key
func SpaceInvaders(w http.ResponseWriter, r *http.Request) {
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
	if isCached(&w, r, key, colors, size) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	write.ImageSVG(w)
	spaceinvaders.SpaceInvaders(w, key, colors, size)
}

func isCached(w *http.ResponseWriter, r *http.Request, key string, colors []color.RGBA, size int) bool {
	e := `"` + key + tgColors.ArrayToHexString(colors) + fmt.Sprintf("%d", size) + `"`
	(*w).Header().Set("Etag", e)
	(*w).Header().Set("Cache-Control", "max-age=2592000") // 30 days
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			return true
		}
	}
	return false
}
