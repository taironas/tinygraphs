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

func LegsFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return int(val%3) + 2
	}
	return 3
}

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

func AnthenasFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return int(val % 4)
	}
	return 2
}

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

func HasFootFromKey(c uint8) bool {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return val%2 == 0
	}
	return true
}

func HasArmsUpFromKey(c uint8) bool {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		return val%2 == 0
	}
	return true
}

func ArmSizeFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val%2 != 0 {
			return 3
		}
	}
	return 2
}

func AnthenaSizeFromKey(c uint8) int {
	s := hex.EncodeToString([]byte{c})
	if val, err := strconv.ParseInt(s, 16, 0); err == nil {
		if val%2 == 0 {
			return 1
		}
	}
	return 2
}
