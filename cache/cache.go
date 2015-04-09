package Cache

import (
	"fmt"
	"image/color"
	"net/http"
	"strings"

	tgColors "github.com/taironas/tinygraphs/colors"
)

func IsCached(w *http.ResponseWriter, r *http.Request, key string, colors []color.RGBA, size int) bool {
	e := `"` + key + tgColors.ArrayToHexString(colors) + fmt.Sprintf("%d", size) + `"`
	(*w).Header().Set("Etag", e)
	(*w).Header().Set("Cache-Control", "max-age=2592000") // 30 days
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			return true
		}
	}
	return false
}
