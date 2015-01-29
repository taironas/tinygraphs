package draw

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strconv"
)

// RandomColorFromArray returns a random color from the given array.
func RandomColorFromArray(colors []color.RGBA) color.RGBA {
	r := rand.Intn(len(colors))
	return colors[r]
}

// FillFromRGBA return a "fill" SVG style from a color.RGBA
func FillFromRGBA(c color.RGBA) string {
	return fmt.Sprintf("fill:rgb(%d,%d,%d)", c.R, c.G, c.B)
}

// PickColor returns a color given a key string, an array of colors and an index.
// key: should be a md5 hash string.
// index: is an index from the key string. Should be in interval [0, 16]
// Algorithm: PickColor converts the key[index] value to a decimal value.
// We pick the ith colors that respects the equality value%numberOfColors == i.
func PickColor(key string, colors []color.RGBA, index int) color.RGBA {
	n := len(colors)
	s := hex.EncodeToString([]byte{key[index]})
	if r, err := strconv.ParseInt(s, 16, 0); err == nil {
		for i := 0; i < n; i++ {
			if int(r)%n == i {
				return colors[i]
			}
		}
	} else {
		log.Println("Error calling ParseInt(%v, 16, 0)", s, err)
	}
	return colors[0]
}
