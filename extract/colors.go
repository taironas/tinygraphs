package extract

import (
	"fmt"
	"image/color"
	"net/http"
	"strconv"
	"strings"

	tgColors "github.com/taironas/tinygraphs/colors"
)

// Colors returns an array of colors based on a HTTP request.
// It follows some rules:
// - if the URL has it's own colors using the 'colors' parameter we return those.
//   e.g: some/url?colors=FFFFFF&colors=222222
// - if the URL has a theme defined using the 'theme' parameter, we return those colors.
//      if the theme does not exist we use the base theme.
//      if the number of colors is specified usign the 'numcolors' parameter,
//      we only return the number of colors specified.
//   e.g: some/url?theme=frogideas
//   returns an array with the first 2 colors that define the 'frogideas' theme.
//   e.g: some/url?theme=frogideas&numcolors=4
//   returns an array with the first 4 colors that define the 'frogideas' theme.
// - if the URL has a background and/or a foreground defined by usign parameters 'bg' and 'fg'
//     we return those.
//   e.g: some/url?bg=FFFFFF&fg=222222
func Colors(r *http.Request) (colors []color.RGBA) {

	if newColors, err := UserColors(r); err == nil {
		return newColors
	}

	th := Theme(r)

	if th == "base" {
		bg, fg := ExtraColors(r)
		colors = append(colors, bg, fg)
	} else {
		m := tgColors.MapOfColorThemes()
		if _, ok := m[th]; ok {
			n := NumColors(r)
			colors = append(colors, m[th][0:n]...)
		} else {
			colors = append(colors, m["base"]...)
		}
	}
	return colors
}

func GColors(r *http.Request) (gColors []color.RGBA) {

	theme := Theme(r)
	if theme != "base" {
		colorMap := tgColors.MapOfColorThemes()
		if _, ok := colorMap[theme]; ok {
			return colorMap[theme][1:3]
		}
	}

	if newColors, err := UserColors(r); err == nil {
		if len(newColors) > 2 {
			gColors = newColors[1:]
		} else {
			gColors = newColors
		}
	} else {
		gColors = Colors(r)
	}

	return
}

// Background extract hexadecimal code background from HTTP request and return color.RGBA
func Background(req *http.Request) (color.RGBA, error) {
	bg := req.FormValue("bg")
	if len(bg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: empty input")
	}
	if r, g, b, err := hexToRGB(bg); err != nil {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	} else {
		return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, nil
	}
}

// Foreground extract hexadecimal code foreground from HTTP request and return color.RGBA
func Foreground(req *http.Request) (color.RGBA, error) {
	fg := req.FormValue("fg")
	if len(fg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: empty input")
	}
	if r, g, b, err := hexToRGB(fg); err != nil {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	} else {
		return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, nil
	}
}

// ExtraColors returns a background and foreground color.RGBA is specified.
// returns black and white otherwise
func ExtraColors(req *http.Request) (color.RGBA, color.RGBA) {
	var err error
	var bg, fg color.RGBA
	if bg, err = Background(req); err != nil {
		bg = tgColors.White()
	}

	if fg, err = Foreground(req); err != nil {
		fg = tgColors.Black()
	}
	return bg, fg
}

// UserColors extract an array of hexadecimal colors and returns an array of color.RGBA
func UserColors(req *http.Request) ([]color.RGBA, error) {
	if err := req.ParseForm(); err != nil {
		return []color.RGBA{}, fmt.Errorf("colors: unable to parse form")
	}

	var colors []color.RGBA
	strColors := req.Form["colors"]
	if len(strColors) == 0 {
		return []color.RGBA{}, fmt.Errorf("colors: empty input")
	}

	for _, c := range strColors {
		if r, g, b, err := hexToRGB(c); err != nil {
			return []color.RGBA{}, fmt.Errorf("colors: wrong input")
		} else {
			new := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}
			colors = append(colors, new)
		}
	}
	return colors, nil
}

// NumColors returns the number of colors in http request.
// Right now we support numbers between 2 and 4.
// Default value 2.
func NumColors(r *http.Request) int64 {
	s := strings.ToLower(r.FormValue("numcolors"))
	if len(s) > 0 {
		if n, err := strconv.ParseInt(s, 0, 64); err == nil {
			if n >= 2 && n <= 4 {
				return n
			}
		}
	}
	return 2
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
