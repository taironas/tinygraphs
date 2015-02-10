package spaceinvaders

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

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
			fill := draw.FillFromRGBA(colors[0])
			if _, ok := colorMap[xQ]; !ok {
				if float64(xQ) < middle {
					colorMap[xQ] = draw.PickColor(key, colors, xQ+2*yQ)
				} else if xQ < squares {
					colorMap[xQ] = colorMap[squares-xQ-1]
				} else {
					colorMap[xQ] = colorMap[0]
				}
			}

			highBodyIndex := 2
			if invader.height > 7 {
				if yQ == highBodyIndex {
					if invader.eyes == 2 {
						if xQ > 4 && xQ < 6 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else if invader.eyes == 3 {
						if xQ > 3 && xQ < 7 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else if invader.eyes == 4 {
						if xQ > 3 && xQ < 8 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
					if invader.eyes > 1 {
						highBodyIndex--
					}
				}
			}

			if yQ == highBodyIndex-1 && invader.anthenaSize == 2 {
				if invader.anthenas == 1 {
					if xQ == 5 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.anthenas == 2 {
					if xQ == 3 || xQ == 7 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.anthenas == 3 {
					if xQ == 2 || xQ == 5 || xQ == 8 {
						fill = draw.FillFromRGBA(colors[1])
					}
				}
			}

			if yQ == highBodyIndex {
				if invader.anthenas == 1 {
					if xQ == 5 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.anthenas == 2 {
					if xQ == 4 || xQ == 6 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.anthenas == 3 {
					if xQ == 3 || xQ == 5 || xQ == 7 {
						fill = draw.FillFromRGBA(colors[1])
					}
				}
			}

			if yQ == 3 { // pre frontal lobe :p
				if invader.eyes == 1 {
					if xQ >= 5 && xQ <= 5 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 2 {
					if xQ >= 4 && xQ <= 6 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 3 {
					if xQ >= 3 && xQ <= 7 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 4 {
					if xQ >= 3 && xQ <= 8 {
						fill = draw.FillFromRGBA(colors[1])
					}
				}
			}

			if yQ == 4 { // frontal lobe
				if invader.eyes == 1 {
					if xQ >= 4 && xQ <= 6 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 2 {
					if xQ >= 3 && xQ <= 7 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 3 {
					if xQ >= 2 && xQ <= 8 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 4 {
					if xQ >= 2 && xQ <= 9 {
						fill = draw.FillFromRGBA(colors[1])
					}
				}
			}

			if yQ == 5 {
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if xQ == (leftOver/2) || xQ == squares-1-leftOver/2 || xQ == (leftOver/2)-1 || xQ == (squares-leftOver/2) {
						fill = draw.FillFromRGBA(colors[1])
					}
				}
			}
			if yQ == 4 || yQ == 6 { // arm extension
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if yQ == 4 && invader.armsUp && invader.armSize == 3 {
						if xQ == (leftOver/2)-1 || xQ == squares-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
					if yQ == 6 && !invader.armsUp && invader.armSize == 3 {
						if xQ == (leftOver/2)-1 || xQ == squares-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}
			}

			if yQ == 5 { // eyes
				if invader.eyes == 1 {
					if xQ == 5 {
						fill = draw.FillFromRGBA(colors[0])
					} else if xQ == 4 || xQ == 6 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 2 {
					if xQ == 4 || xQ == 6 {
						fill = fillWhite()
					} else if xQ == 3 || xQ == 5 || xQ == 7 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 3 {
					if xQ == 5 || xQ == 3 || xQ == 7 {
						fill = draw.FillFromRGBA(colors[0])
					} else if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
						fill = draw.FillFromRGBA(colors[1])
					}
				} else if invader.eyes == 4 {
					if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
						fill = draw.FillFromRGBA(colors[0])
					} else if xQ == 1 || xQ == 3 || xQ == 5 || xQ == 7 || xQ == 9 {
						fill = draw.FillFromRGBA(colors[1])
					}
				}
			}

			if yQ == 6 { // length of body
				leftOver := squares - invader.length
				if xQ > (leftOver/2)-1 && xQ < squares-leftOver/2 {
					fill = draw.FillFromRGBA(colors[1])
				}
			}

			lowBodyIndex := 7
			if invader.height > 5 {
				// add more body if height > 6
				if yQ == lowBodyIndex {
					leftOver := squares - invader.length
					if xQ > (leftOver/2) && xQ < (squares-1-leftOver/2) {
						fill = draw.FillFromRGBA(colors[1])
					}

				}
				lowBodyIndex++
			}

			if invader.height > 6 {
				// add more body if height > 6
				if yQ == lowBodyIndex {
					leftOver := squares - invader.length
					if xQ > (leftOver/2)+1 && xQ < (squares-2-leftOver/2) {
						fill = draw.FillFromRGBA(colors[1])
					}

				}
				lowBodyIndex++
			}

			if yQ == 4 && invader.armsUp && invader.armSize == 3 {
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if (squares - (leftOver / 2) - (leftOver / 2) - 1) >= invader.length {
						if xQ == (leftOver/2)-2 || xQ == squares-leftOver/2+1 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else {
						if xQ == (leftOver/2)-1 || xQ == squares-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}

					}
				}
			}

			// arm up extension.
			if yQ == 4 && invader.armsUp {
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if (squares - (leftOver / 2) - (leftOver / 2) - 1) >= invader.length {
						if xQ == (leftOver/2)-1 || xQ == squares-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else {
						if xQ == (leftOver/2)-2 || xQ == squares+1-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}
			}

			armIndex := 0
			if invader.height > 6 {
				armIndex = lowBodyIndex - 3
			} else {
				armIndex = lowBodyIndex - 2
			}
			if yQ == armIndex && !invader.armsUp && invader.armSize < 3 {
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if (squares - leftOver/2 - (leftOver / 2) - 1) >= invader.length {
						if xQ == (leftOver/2)-1 || xQ == squares-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else {
						if xQ == (leftOver/2)-2 || xQ == squares+1-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}
			}

			// arm down extension.
			if yQ == armIndex && !invader.armsUp && invader.armSize == 3 {
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if (squares - leftOver/2 - (leftOver / 2) - 1) >= invader.length {
						if xQ == (leftOver/2)-1 || xQ == squares-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else {
						if xQ == (leftOver/2)-2 || xQ == squares+1-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}
			}

			// big arm extension
			if yQ == armIndex+1 && !invader.armsUp && invader.armSize == 3 {
				leftOver := squares - invader.length
				if invader.arms > 0 {
					if (squares - leftOver/2 - (leftOver / 2) - 1) >= invader.length {
						if xQ == (leftOver/2)-2 || xQ == squares-leftOver/2+1 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else {
						if xQ == (leftOver/2)-2 || xQ == squares+1-leftOver/2 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}
			}

			if yQ == lowBodyIndex || yQ == lowBodyIndex+1 { // legs
				if invader.legs%2 == 0 {
					if invader.legs == 2 {
						if xQ == 4 || xQ == 6 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else if invader.legs == 4 {
						if invader.length >= 6 {
							if invader.height >= 7 {
								if yQ == lowBodyIndex {
									if xQ == 3 || xQ == 4 || xQ == 6 || xQ == 7 {
										fill = draw.FillFromRGBA(colors[1])
									}
								} else {
									if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
										fill = draw.FillFromRGBA(colors[1])
									}
								}
							} else {
								if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
									fill = draw.FillFromRGBA(colors[1])
								}
							}
						} else {
							if yQ == lowBodyIndex {
								if xQ == 3 || xQ == 5 || xQ == 5 || xQ == 7 {
									fill = draw.FillFromRGBA(colors[1])
								}

							} else {
								if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
									fill = draw.FillFromRGBA(colors[1])
								}
							}
						}
					}
				} else {
					if invader.legs == 1 {
						if xQ == 5 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else if invader.legs == 3 {
						if invader.length > 5 {
							if xQ == 3 || xQ == 5 || xQ == 7 {
								fill = draw.FillFromRGBA(colors[1])
							}
						} else {
							if yQ == lowBodyIndex {
								if xQ == 4 || xQ == 5 || xQ == 6 {
									fill = draw.FillFromRGBA(colors[1])
								}
							} else {
								if xQ == 3 || xQ == 5 || xQ == 7 {
									fill = draw.FillFromRGBA(colors[1])
								}
							}
						}
					} else if invader.legs == 5 {
						if xQ == 1 || xQ == 3 || xQ == 5 || xQ == 7 || xQ == 9 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}
			}

			if yQ == lowBodyIndex+1 && invader.foot {
				if invader.legs%2 == 0 {
					if invader.legs == 2 {
						if xQ == 3 || xQ == 7 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else if invader.legs == 4 {
						if xQ == 1 || xQ == 9 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				} else {
					if invader.legs == 1 {
						if xQ == 4 || xQ == 6 {
							fill = draw.FillFromRGBA(colors[1])
						}
					} else if invader.legs == 3 {
						if invader.length > 5 {
							if xQ == 2 || xQ == 8 {
								fill = draw.FillFromRGBA(colors[1])
							}
						} else {
							if xQ == 3 || xQ == 7 {
								fill = draw.FillFromRGBA(colors[1])
							}
						}

					} else if invader.legs == 5 {
						if xQ == 0 || xQ == 10 {
							fill = draw.FillFromRGBA(colors[1])
						}
					}
				}

			}

			canvas.Rect(x, y, quadrantSize, quadrantSize, fill)
		}
	}
	canvas.End()
}

func fillWhite() string {
	return "stroke:black;stroke-width:2;fill:rgb(255,255,255)"
}

func fillBlack() string {
	return "stroke:black;stroke-width:2;fill:rgb(0,0,0)"
}

func fillGrey() string {
	return "stroke:black;stroke-width:2;fill:rgb(160,160,160)"
}

func fillGreen() string {
	return "stroke:black;stroke-width:2;fill:rgb(0,255,0)"
}

func fillRed() string {
	return "stroke:black;stroke-width:2;fill:rgb(255,0,0)"
}

func fillBlue() string {
	return "stroke:black;stroke-width:2;fill:rgb(0,0,255)"
}
