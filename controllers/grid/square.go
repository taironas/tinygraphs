package grid

import (
	"crypto/md5"
	"fmt"
	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Square is the handler for /square/[A-Za-z0-9]+/?
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
		bg, err1 := background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := size(r)

		m := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.Square(m, key, bg, fg)
		var img image.Image = m
		write.Image(w, &img)
	}
}

// gridColorHandler is the handler for /square/[0-8]/[a-zA-Z0-9]+/?
// build a 6x6 grid with alternate colors based on the number passed in the url
func SquareColor(w http.ResponseWriter, r *http.Request) {

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

			size := size(r)
			m := image.NewRGBA(image.Rect(0, 0, size, size))
			colorMap := tgColors.MapOfColorPatterns()

			draw.Square(m, key, colorMap[int(colorId)][0], colorMap[int(colorId)][1])
			var img image.Image = m
			write.Image(w, &img)
		} else {
			log.Printf("error when extracting permalink string: %v", err)
		}
	}
}

func background(req *http.Request) (color.RGBA, error) {
	bg := req.FormValue("bg")
	if len(bg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	r, g, b, err := hexToRGB(bg)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, err
}

func foreground(req *http.Request) (color.RGBA, error) {
	fg := req.FormValue("fg")
	if len(fg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	r, g, b, err := hexToRGB(fg)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, err
}

// HexToRGB converts an Hex string to a RGB triple.
func hexToRGB(h string) (uint8, uint8, uint8, error) {
	if len(h) > 0 && h[0] == '#' {
		h = h[1:]
	}
	if len(h) == 3 {
		h = h[:1] + h[:1] + h[1:2] + h[1:2] + h[2:] + h[2:]
	}
	if len(h) == 6 {
		if rgb, err := strconv.ParseUint(string(h), 16, 32); err == nil {
			return uint8(rgb >> 16), uint8((rgb >> 8) & 0xFF), uint8(rgb & 0xFF), nil
		} else {
			return 0, 0, 0, err
		}
	}
	return 0, 0, 0, nil
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
