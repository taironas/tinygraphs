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
	"strconv"
)

// gridColorHandler is the handler for /grid/square/[0-9]+/?
// build a 6x6 grid with alternate colors based on the number passed in the url
func Square(w http.ResponseWriter, r *http.Request) {
	id, err := misc.PermalinkString(r, 3)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		size := size(r)
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		colorMap := colors.MapOfColorPatterns()
		h := md5.New()
		io.WriteString(h, id)
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
			size := size(r)
			m := image.NewRGBA(image.Rect(0, 0, size, size))
			colorMap := colors.MapOfColorPatterns()
			h := md5.New()
			io.WriteString(h, id)
			key := fmt.Sprintf("%x", h.Sum(nil)[:])
			draw.Square(m, key, colorMap[int(colorId)][0], colorMap[int(colorId)][1])
			var img image.Image = m
			write.Image(w, &img)
		} else {
			log.Printf("error when extracting permalink string: %v", err)
		}
	}
}

func size(r *http.Request) int {
	strSize := r.FormValue("size")
	if len(strSize) > 0 {
		if size, errSize := strconv.ParseInt(strSize, 0, 64); errSize == nil {
			isize := int(size)
			if isize > 0 && isize < 1000 {
				return int(size)
			}
		}
	}
	return 240
}
