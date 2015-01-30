package spaceinvaders

import (
	"encoding/hex"
	"fmt"
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
	foot     bool
	arms     int
	armsUp   bool
	anthenas int
	height   int
	lenght   int
	eyes     int
}

func SpaceInvaders(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)
	invader := newInvader(key)
	log.Println(fmt.Sprintf("%+v\n", invader))
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
			if yQ == 2 {
				if invader.anthenas == 1 {
					if xQ == 5 {
						fill = fillBlack()
					}
				} else if invader.anthenas == 2 {
					if xQ == 4 || xQ == 6 {
						fill = fillBlack()
					}
				} else if invader.anthenas == 3 {
					if xQ == 3 || xQ == 5 || xQ == 7 {
						fill = fillBlack()
					}
				}
			}

			if yQ == 3 { // pre frontal lobe :p
				if invader.eyes == 1 {
					if xQ >= 5 && xQ <= 5 {
						fill = fillBlack()
					}
				} else if invader.eyes == 2 {
					if xQ >= 4 && xQ <= 6 {
						fill = fillBlack()
					}
				} else if invader.eyes == 3 {
					if xQ >= 3 && xQ <= 7 {
						fill = fillBlack()
					}
				} else if invader.eyes == 4 {
					if xQ >= 3 && xQ <= 8 {
						fill = fillBlack()
					}
				}
			}
			if yQ == 4 { // frontal lobe
				if invader.eyes == 1 {
					if xQ >= 4 && xQ <= 6 {
						fill = fillBlack()
					}
				} else if invader.eyes == 2 {
					if xQ >= 3 && xQ <= 7 {
						fill = fillBlack()
					}
				} else if invader.eyes == 3 {
					if xQ >= 2 && xQ <= 8 {
						fill = fillBlack()
					}
				} else if invader.eyes == 4 {
					if xQ >= 2 && xQ <= 9 {
						fill = fillBlack()
					}
				}
			}

			if yQ == 5 { // eyes
				if invader.eyes == 1 {
					if xQ == 5 {
						fill = fillWhite()
					} else if xQ == 4 || xQ == 6 {
						fill = fillBlack()

					}
				} else if invader.eyes == 2 {
					if xQ == 4 || xQ == 6 {
						fill = fillWhite()
					} else if xQ == 3 || xQ == 5 || xQ == 7 {
						fill = fillBlack()
					}
				} else if invader.eyes == 3 {
					if xQ == 5 || xQ == 3 || xQ == 7 {
						fill = fillWhite()
					} else if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
						fill = fillBlack()
					}
				} else if invader.eyes == 4 {
					if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
						fill = fillWhite()
					} else if xQ == 1 || xQ == 3 || xQ == 5 || xQ == 7 || xQ == 9 {
						fill = fillBlack()
					}
				}
			}

			if yQ == 6 { // lenght of body and arms
				leftOver := squares - invader.lenght
				if xQ > (leftOver/2)-1 && xQ < squares-leftOver/2 {
					fill = fillBlack()
				} else if invader.arms > 0 {
					if xQ == (leftOver/2) || xQ == (leftOver/2)-1 || xQ == squares-1-leftOver/2 || xQ == squares-leftOver/2 {
						fill = fillBlack()
					}
				}
			}

			if yQ == 7 || yQ == 8 { // legs
				if invader.legs%2 == 0 {
					if invader.legs == 2 {
						if xQ == 4 || xQ == 6 {
							fill = fillBlack()
						}
					} else if invader.legs == 4 {
						if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
							fill = fillBlack()
						}
					}
				} else {
					if invader.legs == 1 {
						if xQ == 5 {
							fill = fillBlack()
						}
					} else if invader.legs == 3 {
						if xQ == 3 || xQ == 5 || xQ == 7 {
							fill = fillBlack()
						}
					} else if invader.legs == 5 {
						if xQ == 1 || xQ == 3 || xQ == 5 || xQ == 7 || xQ == 9 {
							fill = fillBlack()
						}
					}
				}
			}
			if yQ == 8 && invader.foot {
				if invader.legs%2 == 0 {
					if invader.legs == 2 {
						if xQ == 3 || xQ == 7 {
							fill = fillBlack()
						}
					} else if invader.legs == 4 {
						if xQ == 1 || xQ == 9 {
							fill = fillBlack()
						}
					}
				} else {
					if invader.legs == 1 {
						if xQ == 4 || xQ == 6 {
							fill = fillBlack()
						}
					} else if invader.legs == 3 {
						if xQ == 2 || xQ == 8 {
							fill = fillBlack()
						}
					} else if invader.legs == 5 {
						if xQ == 0 || xQ == 10 {
							fill = fillBlack()
						}
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
		invader.arms = int(val % 3)
		if invader.arms%2 != 0 {
			invader.arms++
		}
	} else {
		invader.arms = 2
	}

	s = hex.EncodeToString([]byte{key[2]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		invader.anthenas = int(val % 4)
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
		invader.lenght = int(val%4) + 2
		if invader.lenght <= 3 {
			invader.lenght *= 2
		}
		if invader.lenght%2 == 0 {
			invader.lenght += 1
		}
	} else {
		invader.lenght = 6
	}

	s = hex.EncodeToString([]byte{key[5]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val > 5 && val < 1 {
			invader.eyes = 2
		} else {
			invader.eyes = int(val%3) + 1
		}
	} else {
		invader.eyes = 2
	}

	s = hex.EncodeToString([]byte{key[6]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val%2 == 0 {
			invader.foot = true
		} else {
			invader.foot = false
		}
	} else {
		invader.foot = true
	}

	s = hex.EncodeToString([]byte{key[7]})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val%2 == 0 {
			invader.armsUp = true
		} else {
			invader.armsUp = false
		}
	} else {
		invader.foot = true
	}

	return invader
}

func fillWhite() string {
	return "stroke:black;stroke-width:2;fill:rgb(255,255,255)"
}

func fillBlack() string {
	return "stroke:black;stroke-width:2;fill:rgb(0,0,0)"
}
