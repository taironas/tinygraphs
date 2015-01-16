package squares

import (
	"fmt"
	"image/color"
	"net/http"
	"strconv"
	"strings"
)

// extract hexadecimal code background from HTTP request and return color.RGBA
func background(req *http.Request) (color.RGBA, error) {
	bg := req.FormValue("bg")
	if len(bg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	r, g, b, err := hexToRGB(bg)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, err
}

// extract hexadecimal code foreground from HTTP request and return color.RGBA
func foreground(req *http.Request) (color.RGBA, error) {
	fg := req.FormValue("fg")
	if len(fg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	r, g, b, err := hexToRGB(fg)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, err
}

// extract size from HTTP request and return it.
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
	return 210
}

type Format int

const (
	JPEG Format = iota
	SVG
)

func format(r *http.Request) Format {
	strFmt := strings.ToLower(r.FormValue("fmt"))
	if len(strFmt) > 0 {
		if strFmt == "svg" {
			return SVG
		} else if strFmt == "jpeg" || strFmt == "jpg" {
			return JPEG
		}
	}
	return JPEG
}
