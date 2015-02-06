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

// ColorByPercentage returns a color based on the given percentage and the
// number of colors present in the 'colors' array
func ColorByPercentage(colors []color.RGBA, percentage int) color.RGBA {
	if percentage == 0 {
		percentage = 1
	}
	r := rand.Intn(100)
	cr := rand.Intn(100)

	colorChange := 100 / len(colors)
	frontier := (cr / colorChange)

	if frontier == 0 {
		frontier = 1
	} else if frontier == len(colors) {
		frontier = frontier - 1
	}

	if r < percentage {
		return RandomColorFromArray(colors[0:frontier])
	}
	return RandomColorFromArray(colors[frontier-1:])
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
