package spaceinvaders

import (
	"image/color"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

func SpaceInvaders(w http.ResponseWriter, key string, colors []color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)
	invader := newInvader(key)
	// log.Println(fmt.Sprintf("%+v\n", invader)) // for debug
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
				colorMap[xQ] = selectColor(colorMap, key, colors, middle, xQ, yQ, squares)
			}

			highBodyIndex := 2
			if hasEyeOrAnthena(invader, &highBodyIndex, squares, xQ, yQ) {
				fill = draw.FillFromRGBA(colorMap[xQ])
			}

			if hasArmOrExtension(invader, squares, xQ, yQ) {
				fill = draw.FillFromRGBA(colorMap[xQ])
			}

			if yQ == 5 { // clean eye from arm extension
				if eye, c := hasEye4(invader, colorMap, colors, xQ); eye {
					fill = draw.FillFromRGBA(c)
				}
			}

			if yQ == 6 {
				if hasBody(invader, squares, xQ) {
					fill = draw.FillFromRGBA(colorMap[xQ])
				}
			}

			lowBodyIndex := 7
			if hasBody2(invader, &lowBodyIndex, squares, xQ, yQ) {
				fill = draw.FillFromRGBA(colorMap[xQ])
			}

			if hasArmOrExtension2(invader, lowBodyIndex, squares, xQ, yQ) {
				fill = draw.FillFromRGBA(colorMap[xQ])
			}

			if hasLegOrFoot(invader, lowBodyIndex, xQ, yQ) {
				fill = draw.FillFromRGBA(colorMap[xQ])
			}

			canvas.Rect(x, y, quadrantSize, quadrantSize, fill)
		}
	}
	canvas.End()
}

func selectColor(colorMap map[int]color.RGBA, key string, colors []color.RGBA, middle float64, xQ, yQ, squares int) (c color.RGBA) {
	if float64(xQ) < middle {
		c = draw.PickColor(key, colors[1:], xQ+2*yQ)
	} else if xQ < squares {
		c = colorMap[squares-xQ-1]
	} else {
		c = colorMap[0]
	}
	return
}

func hasEye1(invader invader, xQ int) (eye bool) {
	if invader.eyes == 2 {
		if xQ > 4 && xQ < 6 {
			eye = true
		}
	} else if invader.eyes == 3 {
		if xQ > 3 && xQ < 7 {
			eye = true
		}
	} else if invader.eyes == 4 {
		if xQ > 3 && xQ < 8 {
			eye = true
		}
	}
	return
}

func hasEye2(invader invader, xQ int) (eye bool) {
	if invader.eyes == 1 {
		if xQ >= 5 && xQ <= 5 {
			eye = true
		}
	} else if invader.eyes == 2 {
		if xQ >= 4 && xQ <= 6 {
			eye = true
		}
	} else if invader.eyes == 3 {
		if xQ >= 3 && xQ <= 7 {
			eye = true
		}
	} else if invader.eyes == 4 {
		if xQ >= 3 && xQ <= 8 {
			eye = true
		}
	}
	return
}

func hasEye3(invader invader, xQ int) (eye bool) {
	if invader.eyes == 1 {
		if xQ >= 4 && xQ <= 6 {
			eye = true
		}
	} else if invader.eyes == 2 {
		if xQ >= 3 && xQ <= 7 {
			eye = true
		}
	} else if invader.eyes == 3 {
		if xQ >= 2 && xQ <= 8 {
			eye = true
		}
	} else if invader.eyes == 4 {
		if xQ >= 2 && xQ <= 9 {
			eye = true
		}
	}
	return
}

func hasEye4(invader invader, colorMap map[int]color.RGBA, colors []color.RGBA, xQ int) (eye bool, c color.RGBA) {
	if invader.eyes == 1 {
		if xQ == 5 {
			eye = true
			c = colors[0]
		} else if xQ == 4 || xQ == 6 {
			eye = true
			c = colorMap[xQ]
		}
	} else if invader.eyes == 2 {
		if isEyeFillForTwoEyes(xQ) {
			eye = true
			c = colors[0]
		} else if isEyeBorderForTwoEyes(xQ) {
			eye = true
			c = colorMap[xQ]
		}
	} else if invader.eyes == 3 {
		if isEyeFillForThreeEyes(xQ) {
			eye = true
			c = colors[0]
		} else if isEyeBorderForThreeEyes(xQ) {
			eye = true
			c = colorMap[xQ]
		}
	} else if invader.eyes == 4 {
		if isEyeFillForFourEyes(xQ) {
			eye = true
			c = colors[0]
		} else if isEyeBorderForFourEyes(xQ) {
			eye = true
			c = colorMap[xQ]
		}
	}
	return
}

func isEyeFillForTwoEyes(xQ int) bool {
	return xQ == 4 || xQ == 6
}

func isEyeBorderForTwoEyes(xQ int) bool {
	return xQ == 3 || xQ == 5 || xQ == 7
}

func isEyeFillForThreeEyes(xQ int) bool {
	return xQ == 5 || xQ == 3 || xQ == 7
}

func isEyeBorderForThreeEyes(xQ int) bool {
	return xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8
}

func isEyeFillForFourEyes(xQ int) bool {
	return xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8
}

func isEyeBorderForFourEyes(xQ int) bool {
	return xQ == 1 || xQ == 3 || xQ == 5 || xQ == 7 || xQ == 9
}

func hasAnthenas1(invader invader, xQ int) (anthena bool) {
	if invader.anthenas == 1 {
		if xQ == 5 {
			anthena = true
		}
	} else if invader.anthenas == 2 {
		if xQ == 3 || xQ == 7 {
			anthena = true
		}
	} else if invader.anthenas == 3 {
		if xQ == 2 || xQ == 5 || xQ == 8 {
			anthena = true
		}
	}
	return
}

func hasAnthenas2(invader invader, xQ int) (anthena bool) {
	if invader.anthenas == 1 {
		if xQ == 5 {
			anthena = true
		}
	} else if invader.anthenas == 2 {
		if xQ == 4 || xQ == 6 {
			anthena = true
		}
	} else if invader.anthenas == 3 {
		if xQ == 3 || xQ == 5 || xQ == 7 {
			anthena = true
		}
	}
	return
}

func hasEyeOrAnthena(invader invader, highBodyIndex *int, squares, xQ, yQ int) (result bool) {
	if invader.height > 7 {
		if yQ == *highBodyIndex {
			if hasEye1(invader, xQ) {
				result = true
			}
			if invader.eyes > 1 {
				*highBodyIndex--
			}
		}
	}

	if yQ == *highBodyIndex-1 && invader.anthenaSize == 2 {
		if hasAnthenas1(invader, xQ) {
			result = true
		}
	}

	if yQ == *highBodyIndex {
		if hasAnthenas2(invader, xQ) {
			result = true
		}
	}

	if yQ == 3 {
		if hasEye2(invader, xQ) {
			result = true
		}
	}

	if yQ == 4 {
		if hasEye3(invader, xQ) {
			result = true
		}
	}

	return
}

func hasArm(invader invader, squares, xQ int) (arm bool) {
	if invader.arms <= 0 {
		return
	}
	leftOver := squares - invader.length
	half := leftOver / 2
	if xQ == half || xQ == squares-1-half || xQ == half-1 || xQ == squares-half {
		arm = true
	}
	return
}

func hasArm2(invader invader, squares, xQ, yQ int) (arm bool) {
	if yQ != 4 || !invader.armsUp || invader.armSize != 3 {
		return
	}

	if invader.arms <= 0 {
		return
	}

	leftOver := squares - invader.length
	half := leftOver / 2

	if (squares - half - half - 1) >= invader.length {
		if xQ == half-2 || xQ == squares-half+1 {
			arm = true
		}
	} else {
		if xQ == half-1 || xQ == squares-half {
			arm = true
		}
	}
	return
}

func hasArmExtension(invader invader, squares, xQ, yQ int) (armExtension bool) {
	if invader.arms <= 0 {
		return
	}

	leftOver := squares - invader.length
	half := leftOver / 2

	if yQ == 4 && invader.armsUp && invader.armSize == 3 {
		if xQ == half-1 || xQ == squares-half {
			armExtension = true
		}
	}
	if yQ == 6 && !invader.armsUp && invader.armSize == 3 {
		if xQ == half-1 || xQ == squares-half {
			armExtension = true
		}
	}
	return
}

func hasArmExtension2(invader invader, squares, xQ, yQ int) (armExtension bool) {
	if yQ != 4 || !invader.armsUp {
		return
	}
	if invader.arms <= 0 {
		return
	}

	leftOver := squares - invader.length
	half := leftOver / 2
	if (squares - half - half - 1) >= invader.length {
		if xQ == half-1 || xQ == squares-half {
			armExtension = true
		}
	} else {
		if xQ == half-2 || xQ == squares+1-half {
			armExtension = true
		}
	}
	return
}

func hasArmDown(invader invader, armIndex, squares, xQ, yQ int) (armDown bool) {
	if yQ != armIndex || invader.armsUp || invader.armSize >= 3 {
		return
	}
	if invader.arms <= 0 {
		return
	}

	leftOver := squares - invader.length
	half := leftOver / 2
	if (squares - half - half - 1) >= invader.length {
		if xQ == half-1 || xQ == squares-half {
			armDown = true
		}
	} else {
		if xQ == half-2 || xQ == squares+1-half {
			armDown = true
		}
	}
	return
}

func hasArmDownExtension(invader invader, armIndex, squares, xQ, yQ int) (armDownExtension bool) {
	if yQ != armIndex || invader.armsUp || invader.armSize != 3 {
		return
	}
	if invader.arms <= 0 {
		return
	}

	leftOver := squares - invader.length
	half := leftOver / 2

	if (squares - half - half - 1) >= invader.length {
		if xQ == half-1 || xQ == squares-half {
			armDownExtension = true
		}
	} else {
		if xQ == half-2 || xQ == squares+1-half {
			armDownExtension = true
		}
	}
	return
}

func hasBigArmExtension(invader invader, armIndex, squares, xQ, yQ int) (bigArmExtension bool) {
	if yQ != armIndex+1 || invader.armsUp || invader.armSize != 3 {
		return
	}

	if invader.arms <= 0 {
		return
	}

	leftOver := squares - invader.length
	half := leftOver / 2

	if (squares - half - half - 1) >= invader.length {
		if xQ == half-2 || xQ == squares-half+1 {
			bigArmExtension = true
		}
	} else {
		if xQ == half-2 || xQ == squares+1-half {
			bigArmExtension = true
		}
	}
	return
}

func getArmIndex(invader invader, lowBodyIndex int) (armIndex int) {
	if invader.height > 6 {
		armIndex = lowBodyIndex - 3
	} else {
		armIndex = lowBodyIndex - 2
	}
	return
}

func hasArmOrExtension(invader invader, squares, xQ, yQ int) (result bool) {
	if yQ == 5 {
		if hasArm(invader, squares, xQ) {
			result = true
		}
	}

	if yQ == 4 || yQ == 6 {
		if hasArmExtension(invader, squares, xQ, yQ) {
			result = true
		}
	}
	return
}

func hasArmOrExtension2(invader invader, lowBodyIndex, squares, xQ, yQ int) (result bool) {
	if hasArm2(invader, squares, xQ, yQ) {
		result = true
	}

	if hasArmExtension2(invader, squares, xQ, yQ) {
		result = true
	}

	armIndex := getArmIndex(invader, lowBodyIndex)

	if hasArmDown(invader, armIndex, squares, xQ, yQ) {
		result = true
	}

	if hasArmDownExtension(invader, armIndex, squares, xQ, yQ) {
		result = true
	}

	if hasBigArmExtension(invader, armIndex, squares, xQ, yQ) {
		result = true
	}
	return
}

func hasBody(invader invader, squares, xQ int) (body bool) {
	leftOver := squares - invader.length
	half := leftOver / 2
	if xQ > half-1 && xQ < squares-half {
		body = true
	}
	return
}

func hasBody2(invader invader, lowBodyIndex *int, squares, xQ, yQ int) (result bool) {
	if invader.height > 5 {
		if hasLowBody(invader, squares, *lowBodyIndex, xQ, yQ) {
			result = true
		}
		*lowBodyIndex++
	}

	if invader.height > 6 {
		if hasLowBody2(invader, squares, *lowBodyIndex, xQ, yQ) {
			result = true
		}
		*lowBodyIndex++
	}
	return result

}

func hasLowBody(invader invader, squares, lowBodyIndex, xQ, yQ int) (lowbody bool) {
	if yQ != lowBodyIndex {
		return
	}
	leftOver := squares - invader.length
	half := leftOver / 2
	if xQ > half && xQ < (squares-1-half) {
		lowbody = true
	}
	return
}

func hasLowBody2(invader invader, squares, lowBodyIndex, xQ, yQ int) (lowbody bool) {
	if yQ != lowBodyIndex {
		return
	}
	leftOver := squares - invader.length
	half := leftOver / 2
	if xQ > half+1 && xQ < (squares-2-half) {
		lowbody = true
	}
	return
}

func isOneOfTheTwoLegs(xQ int) bool {
	if xQ == 4 || xQ == 6 {
		return true
	}
	return false
}

func isOneOfTheFourLegsInner(xQ int) bool {
	if xQ == 3 || xQ == 4 || xQ == 6 || xQ == 7 {
		return true
	}
	return false
}

func isOneOfTheFourLegsInnerLow(xQ int) bool {
	if xQ == 3 || xQ == 5 || xQ == 5 || xQ == 7 {
		return true
	}
	return false
}

func isOneOfTheFourLegsOuter(xQ int) bool {
	if xQ == 2 || xQ == 4 || xQ == 6 || xQ == 8 {
		return true
	}
	return false
}

func isOneLeg(xQ int) bool {
	if xQ == 5 {
		return true
	}
	return false
}

func isOneOfTheThreeLegs(invader invader, lowBodyIndex, xQ, yQ int) bool {
	if invader.length > 5 {
		if xQ == 3 || xQ == 5 || xQ == 7 {
			return true
		}
	}

	if yQ == lowBodyIndex {
		if xQ == 4 || xQ == 5 || xQ == 6 {
			return true
		}
	} else {
		if xQ == 3 || xQ == 5 || xQ == 7 {
			return true
		}
	}
	return false
}

func isOneOfTheFourLegs(invader invader, lowBodyIndex, xQ, yQ int) (leg bool) {
	if invader.length >= 6 {
		if invader.height >= 7 {
			if yQ == lowBodyIndex {
				if isOneOfTheFourLegsInner(xQ) {
					leg = true
				}
			} else {
				if isOneOfTheFourLegsOuter(xQ) {
					leg = true
				}
			}
		} else {
			if isOneOfTheFourLegsOuter(xQ) {
				leg = true
			}
		}
	} else {
		if yQ == lowBodyIndex {
			if isOneOfTheFourLegsInnerLow(xQ) {
				leg = true
			}
		} else {
			if isOneOfTheFourLegsOuter(xQ) {
				leg = true
			}
		}
	}
	return
}

func isOneOfTheFiveLegs(xQ int) bool {
	if xQ == 1 || xQ == 3 || xQ == 5 || xQ == 7 || xQ == 9 {
		return true
	}
	return false
}

func hasLeg(invader invader, lowBodyIndex, xQ, yQ int) (leg bool) {
	if invader.legs%2 == 0 {
		if invader.legs == 2 {
			if isOneOfTheTwoLegs(xQ) {
				leg = true
			}
		} else if invader.legs == 4 {
			if isOneOfTheFourLegs(invader, lowBodyIndex, xQ, yQ) {
				leg = true
			}
		}
	} else {
		if invader.legs == 1 {
			if isOneLeg(xQ) {
				leg = true
			}
		} else if invader.legs == 3 {
			if isOneOfTheThreeLegs(invader, lowBodyIndex, xQ, yQ) {
				leg = true
			}
		} else if invader.legs == 5 {
			if isOneOfTheFiveLegs(xQ) {
				leg = true
			}
		}
	}
	return
}

func isOneOfTheTwoFeet(xQ int) bool {
	if xQ == 3 || xQ == 7 {
		return true
	}
	return false
}

func isOneOfTheFourFeet(xQ int) bool {
	if xQ == 1 || xQ == 9 {
		return true
	}
	return false
}

func isOneFoot(xQ int) bool {
	if xQ == 4 || xQ == 6 {
		return true
	}
	return false
}

func isOneOfTheThreeFeet(invader invader, xQ int) bool {
	if invader.length > 5 {
		if xQ == 2 || xQ == 8 {
			return true
		}
	} else {
		if xQ == 3 || xQ == 7 {
			return true
		}
	}
	return false
}

func isOneOfTheFiveFeet(xQ int) bool {
	if xQ == 0 || xQ == 10 {
		return true
	}
	return false
}

func hasFoot(invader invader, lowBodyIndex, xQ, yQ int) (foot bool) {
	if yQ != lowBodyIndex+1 || !invader.foot {
		return
	}

	if invader.legs%2 == 0 {
		if invader.legs == 2 {
			if isOneOfTheTwoFeet(xQ) {
				foot = true
			}
		} else if invader.legs == 4 {
			if isOneOfTheFourFeet(xQ) {
				foot = true
			}
		}
	} else {
		if invader.legs == 1 {
			if isOneFoot(xQ) {
				foot = true
			}
		} else if invader.legs == 3 {
			if isOneOfTheThreeFeet(invader, xQ) {
				foot = true
			}
		} else if invader.legs == 5 {
			if isOneOfTheFiveFeet(xQ) {
				foot = true
			}
		}
	}

	return
}

func hasLegOrFoot(invader invader, lowBodyIndex, xQ, yQ int) bool {
	if yQ == lowBodyIndex || yQ == lowBodyIndex+1 {
		if hasLeg(invader, lowBodyIndex, xQ, yQ) {
			return true
		}
	}

	if hasFoot(invader, lowBodyIndex, xQ, yQ) {
		return true
	}
	return false
}
