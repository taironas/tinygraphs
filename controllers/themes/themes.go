package themes

import (
	"log"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/colors"
)

// Theme handler builds an image with the colors of a theme
// the theme is defined by keyword :theme
// url: "/themes/:theme"
func Theme(w http.ResponseWriter, r *http.Request) {

	var err error
	var theme string
	if theme, _ = route.Context.Get(r, "theme"); err != nil {
		theme = "base"
	}

	colorMap := colors.MapOfColorThemes()
	if val, ok := colorMap[theme]; ok {
		log.Println(val)
	}
}
