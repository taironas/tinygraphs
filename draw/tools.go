package draw

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
)

// RandomColorFromArray returns a random color from the given array.
func RandomColorFromArray(colors []color.RGBA) color.RGBA {
	r := rand.Intn(len(colors))
	return colors[r]
}

// RandomIndexFromArray returns an index from the given array.
func RandomIndexFromArray(colors []color.RGBA) int {
	r := rand.Intn(len(colors))
	return r
}

// ColorByPercentage returns a color based on the given percentage and the
// number of colors present in the 'colors' array in a gradient way.
func ColorByPercentage(colors []color.RGBA, percentage int) color.RGBA {
	r := rand.Intn(100)
	colorChange := 100 / len(colors)
	frontier := 0
	frontier = int(math.Ceil(float64(r) / float64(colorChange)))

	if r < percentage {
		if frontier == 0 {
			frontier++
		}
		return RandomColorFromArray(colors[:frontier])
	}
	if frontier == len(colors) {
		frontier--
	}
	return RandomColorFromArray(colors[frontier:])
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

func PickIndex(key string, n int, index int) int {
	s := hex.EncodeToString([]byte{key[index]})
	if r, err := strconv.ParseInt(s, 16, 0); err == nil {
		for i := 0; i < n; i++ {
			if int(r)%n == i {
				return i
			}
		}
	}
	return 0
}

// RGBToHex converts an RGB triple to an Hex string.
func RGBToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
