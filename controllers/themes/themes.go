package themes

import (
	"image"
	"image/color"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// Theme handler builds an image with the colors of a theme
// the theme is defined by keyword :theme
// url: "/themes/:theme"
func Theme(w http.ResponseWriter, r *http.Request) {
	var err error
	var th string
	if th, _ = route.Context.Get(r, "theme"); err != nil {
		th = "base"
	}

	colorMap := colors.MapOfColorThemes()
	var theme []color.RGBA
	if val, ok := colorMap[th]; ok {
		theme = val
	} else {
		theme = colorMap["base"]
	}

	width := extract.WidthOrDefault(r, 20*len(theme))
	height := extract.HeightOrDefault(r, 34)

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, width, height))
		squares.Palette(m, theme)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.PaletteSVG(w, theme, width, height)
	}
}
