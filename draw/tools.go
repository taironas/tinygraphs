package draw

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strconv"
)

// randomColor returns a random color between c1 and c2
func randomColor(c1, c2 color.RGBA) color.RGBA {
	r := rand.Intn(2)
	if r == 1 {
		return c1
	}
	return c2
}

// getRandomColor returns a random color between c1 and c2
func randomColorFromArray(colors []color.RGBA) color.RGBA {
	r := rand.Intn(len(colors))
	return colors[r]
}

func fillFromRGBA(c color.RGBA) string {
	return fmt.Sprintf("fill:rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func colorFromKey(key string, color1, color2 color.RGBA, index int) color.RGBA {
	s := hex.EncodeToString([]byte{key[index]})
	if r, err := strconv.ParseInt(s, 16, 0); err == nil {
		if r%2 == 0 {
			return color1
		}
		return color2
	} else {
		log.Println("Error calling ParseInt(%v, 16, 0)", s, err)
	}
	return color1
}

func colorFromKeyAndArray(key string, colors []color.RGBA, index int) color.RGBA {
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
