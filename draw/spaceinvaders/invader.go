package spaceinvaders

import (
	"encoding/hex"
	"strconv"
)

type invader struct {
	legs        int
	foot        bool
	arms        int
	armsUp      bool
	anthenas    int
	height      int
	length      int
	eyes        int
	armSize     int
	anthenaSize int
}

// newInvader returns a invader type built from key which is a 16 md5 hash.
func newInvader(key string) invader {
	invader := invader{}
	invader.legs = LegsFromKey(key[0])
	invader.arms = ArmsFromKey(key[1])
	invader.anthenas = AnthenasFromKey(key[2])
	invader.length = LengthFromKey(key[3])
	invader.height = HeightFromKey(key[5])
	invader.eyes = EyesFromKey(key[6])
	invader.foot = HasFootFromKey(key[7])
	invader.armsUp = HasArmsUpFromKey(key[8])
	invader.armSize = ArmSizeFromKey(key[9])
	invader.anthenaSize = AnthenaSizeFromKey(key[10])
	return invader
}

// LegsFromKey returns the number of legs that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 3.
func LegsFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return int(val%3) + 2
	}
	return 3
}

// ArmsFromKey returns the number of arms that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 2.
func ArmsFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		v := int(val % 3)
		if v%2 != 0 {
			v++
			return v
		}
	}
	return 2
}

// AnthenasFromKey returns the number of anthenas that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 2.
func AnthenasFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return int(val % 4)
	}
	return 2
}

// LengthFromKey returns the number of anthenas that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 7.
func LengthFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		v := int(val % 8)
		if v < 4 {
			return 5
		}
		return v
	}
	return 7
}

// HeightFromKey returns the height that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 6.
func HeightFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		v := int(val % 10)
		if v < 6 {
			v = 6
		}
		return v
	}
	return 6
}

// EyesFromKey returns the number of eyes that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 2.
func EyesFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		var v int
		if val > 5 && val < 1 {
			v = 2
		} else {
			v = int(val%3) + 1
		}
		return v
	}
	return 2
}

// HasFootFromKey tells you if the invader should have feets or not,
// with respect to the character passed as argument.
// If it fails to parse the character it return true.
func HasFootFromKey(c uint8) bool {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return val%2 == 0
	}
	return true
}

// HasArmsUpFromKey tells you if the invader should it's arms up or not,
// with respect to the character passed as argument.
// If it fails to parse the character it return true.
func HasArmsUpFromKey(c uint8) bool {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return val%2 == 0
	}
	return true
}

// ArmSizeFromKey returns the size of the arms that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 2.
func ArmSizeFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val%2 != 0 {
			return 3
		}
	}
	return 2
}

// AnthenaSizeFromKey returns the size of the anthenas that the invader should have,
// with respect to the character passed as argument.
// If it fails to parse the character it return the value of 2.
func AnthenaSizeFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val%2 == 0 {
			return 1
		}
	}
	return 2
}
