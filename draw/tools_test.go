package draw

import (
	"image/color"
	"testing"
)

var (
	c1     = color.RGBA{255, 245, 249, 255}
	colors = []color.RGBA{
		c1,
		color.RGBA{232, 70, 134, 255},
		color.RGBA{232, 70, 186, 255},
		color.RGBA{232, 70, 81, 255},
	}
)

func TestRandomColorFromArrayWithFreq(t *testing.T) {

	if c := RandomColorFromArrayWithFreq(colors, 0); c == c1 {
		t.Errorf("expected color different than", c1, "got", c)
	}

	if c := RandomColorFromArrayWithFreq(colors, 1); c != c1 {
		t.Errorf("expected color", c1, "got", c)
	}
}

func TestRandomColorArray(t *testing.T) {

	if c := RandomColorFromArray(colors); !contains(colors, c) {
		t.Errorf("expected color in array", colors, "got", c)
	}
}

func TestRandomIndexFromArrayWithFreq(t *testing.T) {

	if i := RandomIndexFromArrayWithFreq(colors, 1); i != 0 {
		t.Errorf("expected index 0 got", i)
	}
	if i := RandomIndexFromArrayWithFreq(colors, 0); i < 0 && i >= len(colors) {
		t.Errorf("expected index between 0 and", len(colors), "got", i)
	}
}

func TestRandomIndexFromArray(t *testing.T) {

	if i := RandomIndexFromArray(colors); i < 0 && i >= len(colors) {
		t.Errorf("expected index between 0 and", len(colors), "got", i)
	}
}

func TestColorByPercentage(t *testing.T) {

	if c := ColorByPercentage(colors, 100); !contains(colors, c) {
		t.Errorf("expected color in array ", colors, "got", c)
	}
	if c := ColorByPercentage(colors, 0); c == c1 {
		t.Errorf("expected color equal than ", c1, "got", c)
	}
}

func TestFillFromRGBA(t *testing.T) {

	expected := "fill:rgb(255,245,249)"
	if s := FillFromRGBA(c1); s != expected {
		t.Errorf("expected ", expected, "got", s)
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
