package spaceinvaders

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
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

	colorMap := colors.MapOfColorThemes()
	bg, fg := extract.ExtraColors(r, colorMap)
	theme := extract.Theme(r)
	numColors := extract.NumColors(r)

	var colors []color.RGBA
	if theme != "base" {
		if _, ok := colorMap[theme]; ok {
			colors = append(colors, colorMap[theme][0:numColors]...)
		} else {
			colors = append(colors, colorMap["base"]...)
		}
	} else {
		colors = append(colors, bg, fg)
	}

	size := extract.Size(r)

	write.ImageSVG(w)
	spaceinvaders.SpaceInvaders(w, key, colors, size)
}
