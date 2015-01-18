// Package extract provides a set of functions to extract parameters from an http:Request.
package extract

import (
	"fmt"
	"github.com/taironas/tinygraphs/format"
	"image/color"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// extract hexadecimal code background from HTTP request and return color.RGBA
func Background(req *http.Request) (color.RGBA, error) {
	bg := req.FormValue("bg")
	if len(bg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	r, g, b, err := hexToRGB(bg)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, err
}

// extract hexadecimal code foreground from HTTP request and return color.RGBA
func Foreground(req *http.Request) (color.RGBA, error) {
	fg := req.FormValue("fg")
	if len(fg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	r, g, b, err := hexToRGB(fg)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, err
}

// extract size from HTTP request and return it.
func Size(r *http.Request) int {
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

func Format(r *http.Request) format.Format {
	// strFmt := strings.ToLower(r.FormValue("fmt"))
	// if len(strFmt) > 0 {
	// 	if strFmt == "svg" {
	// 		return format.SVG
	// 	} else if strFmt == "jpeg" || strFmt == "jpg" {
	// 		return format.JPEG
	// 	}
	// }
	strFmt := strings.ToLower(path.Ext(r.URL.Path))
	if strFmt != ".jpeg" && strFmt != ".jpg" && strFmt != ".svg" {
		strFmt = "jpeg"
	}
	if strFmt == ".svg" {
		return format.SVG
	} else if strFmt == ".jpeg" || strFmt == ".jpg" {
		return format.JPEG
	}
	return format.JPEG
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
