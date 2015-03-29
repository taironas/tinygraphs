package draw

import (
	"image/color"
	"testing"
)

func TestRandomColorFromArrayWithFreq(t *testing.T) {

	c1 := color.RGBA{255, 245, 249, 255}
	colors := []color.RGBA{
		c1,
		color.RGBA{232, 70, 134, 255},
		color.RGBA{232, 70, 186, 255},
		color.RGBA{232, 70, 81, 255},
	}

	if c := RandomColorFromArrayWithFreq(colors, 0); c == c1 {
		t.Errorf("expected color different than", c1, "got", c)
	}

	if c := RandomColorFromArrayWithFreq(colors, 1); c != c1 {
		t.Errorf("expected color", c1, "got", c)
	}
}

func TestRandomColorArray(t *testing.T) {

	c1 := color.RGBA{255, 245, 249, 255}
	colors := []color.RGBA{
		c1,
		color.RGBA{232, 70, 134, 255},
		color.RGBA{232, 70, 186, 255},
		color.RGBA{232, 70, 81, 255},
	}

	if c := RandomColorFromArray(colors); !contains(colors, c) {
		t.Errorf("expected color in array", colors, "got", c)
	}
}

func contains(a []color.RGBA, e color.RGBA) bool {
	for _, v := range a {
		if v == e {
			return true
		}
	}
	return false
}
