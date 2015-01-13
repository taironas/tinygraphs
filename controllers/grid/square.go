package grid

import (
	"crypto/md5"
	"fmt"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"io"
	"log"
	"net/http"
)

// gridColorHandler is the handler for /grid/square/[0-9]+/?
// build a 6x6 grid with alternate colors based on the number passed in the url
func Square(w http.ResponseWriter, r *http.Request) {
	id, err := misc.PermalinkString(r, 3)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		colorMap := colors.MapOfColorPatterns()
		h := md5.New()
		io.WriteString(h, id)
		log.Printf("md5: %x", h.Sum(nil))
		key := fmt.Sprintf("%x", h.Sum(nil)[:])
		draw.Square(m, key, colorMap[0][0], colorMap[0][1])
		var img image.Image = m
		write.Image(w, &img)
	}
}

// gridColorHandler is the handler for /grid/square/[0-8]/[a-zA-Z0-9]+/?
// build a 6x6 grid with alternate colors based on the number passed in the url
func SquareColor(w http.ResponseWriter, r *http.Request) {

	if colorId, err := misc.PermalinkID(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		if id, err1 := misc.PermalinkString(r, 4); err1 == nil {
			m := image.NewRGBA(image.Rect(0, 0, 240, 240))
			colorMap := colors.MapOfColorPatterns()
			h := md5.New()
			io.WriteString(h, id)
			log.Printf("md5: %x", h.Sum(nil))
			key := fmt.Sprintf("%x", h.Sum(nil)[:])
			draw.Square(m, key, colorMap[int(colorId)][0], colorMap[int(colorId)][1])
			var img image.Image = m
			write.Image(w, &img)
		} else {
			log.Printf("error when extracting permalink string: %v", err)
		}
	}
}
