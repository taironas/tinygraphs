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
// - if you use parameter 'inv' you can invert the colors.
//   This is true only if the number of colors is equal to 2 and you use 'theme' colors or default ones.
//   e.g: some/url?inv=1
//        some/url?theme=frogideas&numcolors=2&inv=1
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

	if len(colors) == 2 && Inverse(r) {
		swap(&colors)

	}

	if order := Order(r); len(order) > 0 {
		ReOrder(order, &colors)
	}
	return
}

// swap modifies the color array passed as para.
// it will swap element 0 to 1 and 1 to 0
func swap(pColors *[]color.RGBA) {
	colors := *pColors
	tmp := colors[0]
	colors[0] = colors[1]
	colors[1] = tmp
}

// ReOrder will change the order of the color array passed as param
// with respect to the 'order' array of integers.
// both array have to have the same length.
// will ignore any change if the indexes in the order array are out of range.
func ReOrder(order []int, pColors *[]color.RGBA) {
	colors := *pColors
	if len(order) == len(colors) {
		tmp := []color.RGBA{}
		reOrder := true
		for i := range order {
			if order[i] >= 0 && order[i] < len(colors) {
				tmp = append(tmp, colors[order[i]])
			} else {
				reOrder = false
				break
			}
		}
		if reOrder {
			*pColors = tmp
		}
	}
}

// GColors returns an array of colors based on a HTTP request.
// It follows some rules:
// - if the URL has it's a theme defined by using the 'theme' parameter we return
//   all the colors of that theme but the first one.
// - if the URL has it's own colors using the 'colors' parameter we return all of
//   those colors but the first one defined. If the array is less or equal to 2 we return the
//   full array.
//   e.g: some/url?colors=FFFFFF&colors=222222&colors=111111 gives you the
//   following array [111111, 222222]
// - otherwise we return Colors(r)
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
	r, g, b, err := hexToRGB(bg)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, nil
}

// Foreground extract hexadecimal code foreground from HTTP request and return color.RGBA
func Foreground(req *http.Request) (color.RGBA, error) {
	fg := req.FormValue("fg")
	if len(fg) == 0 {
		return color.RGBA{}, fmt.Errorf("background: empty input")
	}
	r, g, b, err := hexToRGB(fg)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("background: wrong input")
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}, nil
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

// UserColors extract an array of hexadecimal colors from an HTTP request
// and returns an array of color.RGBA
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
		r, g, b, err := hexToRGB(c)
		if err != nil {
			return []color.RGBA{}, fmt.Errorf("colors: wrong input")
		}
		new := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}
		colors = append(colors, new)
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
		rgb, err := strconv.ParseUint(string(h), 16, 32)
		if err != nil {
			return 0, 0, 0, err
		}
		return uint8(rgb >> 16), uint8((rgb >> 8) & 0xFF), uint8(rgb & 0xFF), nil
	}
	return 0, 0, 0, nil
}
