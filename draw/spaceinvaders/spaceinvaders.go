package spaceinvaders

import (
	"encoding/hex"
	"image/color"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

type invader struct {
	legs     int
	arms     int
	anthenas int
	height   int
	lenght   int
	eyes     int
}

func SpaceInvaders(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)
	invader := newInvader(key)
	log.Println(invader)
	squares := 11
	quadrantSize := size / squares
	middle := math.Ceil(float64(squares) / float64(2))
	colorMap := make(map[int]color.RGBA)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)

		for xQ := 0; xQ < squares; xQ++ {
			x := xQ * quadrantSize
			fill := fillWhite()
			if _, ok := colorMap[xQ]; !ok {
				if float64(xQ) < middle {
					colorMap[xQ] = draw.PickColor(key, colors, xQ+2*yQ)
				} else if xQ < squares {
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorMap[xQ] = colorMap[0]
				}
			}
			// todo(santiago): needs refactor.
			if yQ == 5 {
				if invader.eyes == 1 {
					if xQ == 5 {
						fill = fillBlack()
					}
				} else if invader.eyes == 2 {
					if xQ == 4 || xQ == 6 {
						fill = fillBlack()
					}
				} else if invader.eyes == 3 {
					if xQ == 5 || xQ == 3 || xQ == 7 {
						fill = fillBlack()
					}
				} else if invader.eyes == 4 {
					if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
						fill = fillBlack()
					}
				}
			}

			canvas.Rect(x, y, quadrantSize, quadrantSize, fill) //draw.FillFromRGBA(colorMap[xQ]))
		}
	}
	canvas.End()
}

// key is a 16 md5 hash
func newInvader(key string) invader {
	invader := invader{}
	var s string

	s = hex.EncodeToString([]byte{key[0]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		invader.legs = int(val%3) + 2
	} else {
		invader.legs = 3
	}

	s = hex.EncodeToString([]byte{key[1]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		invader.arms = int(val%3) + 2
	} else {
		invader.arms = 2
	}

	s = hex.EncodeToString([]byte{key[2]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		invader.anthenas = int(val%2) + 1
	} else {
		invader.anthenas = 2
	}

	s = hex.EncodeToString([]byte{key[3]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val < 7 {
			invader.height = 8
		} else {
			invader.height = 8
		}
	} else {
		invader.height = 8
	}

	s = hex.EncodeToString([]byte{key[4]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		invader.lenght = int(val%5) + 2
	} else {
		invader.lenght = 5
	}

	s = hex.EncodeToString([]byte{key[4]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val > 5 && val < 1 {
			invader.eyes = 2
		} else {
			invader.eyes = int(val%3) + 1
		}
	} else {
		invader.eyes = 2
	}

	return invader
}

func fillWhite() string {
	return "stroke:black;stroke-width:2;fill:rgb(255,255,255)"
}

func fillBlack() string {
	return "stroke:black;stroke-width:2;fill:rgb(0,0,0)"

}
