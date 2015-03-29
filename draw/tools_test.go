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
