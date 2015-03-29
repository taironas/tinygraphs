package draw

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
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
		t.Errorf("expected color different to %v got %v", c1, c)
	}

	if c := RandomColorFromArrayWithFreq(colors, 1); c != c1 {
		t.Errorf("expected color %v got %v", c1, c)
	}
}

func TestRandomColorArray(t *testing.T) {

	if c := RandomColorFromArray(colors); !contains(colors, c) {
		t.Errorf("expected color in array %v got %v", colors, c)
	}
}

func TestRandomIndexFromArrayWithFreq(t *testing.T) {

	if i := RandomIndexFromArrayWithFreq(colors, 1); i != 0 {
		t.Errorf("expected index 0 got", i)
	}
	if i := RandomIndexFromArrayWithFreq(colors, 0); i < 0 && i >= len(colors) {
		t.Errorf("expected index between 0 and %v got %v", len(colors), i)
	}
}

func TestRandomIndexFromArray(t *testing.T) {

	if i := RandomIndexFromArray(colors); i < 0 && i >= len(colors) {
		t.Errorf("expected index between 0 and %v got %v", len(colors), i)
	}
}

func TestColorByPercentage(t *testing.T) {

	if c := ColorByPercentage(colors, 100); !contains(colors, c) {
		t.Errorf("expected color in array %v got %v ", colors, c)
	}
	if c := ColorByPercentage(colors, 0); c == c1 {
		t.Errorf("expected color %v different to %v", c1, c)
	}
}

func TestFillFromRGBA(t *testing.T) {

	expected := "fill:rgb(255,245,249)"
	if s := FillFromRGBA(c1); s != expected {
		t.Errorf("expected %v got %v ", expected, s)
	}
}

func TestPickColor(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "hello")
	key := fmt.Sprintf("%x", h.Sum(nil)[:])

	color1 := PickColor(key, colors, 0)
	color2 := PickColor(key, colors, 0)
	if color1 != color2 {
		t.Errorf("expected %v and %v to be equal", color1, color2)
	}
}

func TestPickIndex(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "hello")
	key := fmt.Sprintf("%x", h.Sum(nil)[:])

	i1 := PickIndex(key, 10, 0)
	i2 := PickIndex(key, 10, 0)
	if i1 != i2 {
		t.Errorf("expected %v and %v to be equal", i1, i2)
	}
}

func TestRGBToHex(t *testing.T) {
	expected := "#FFF5F9"
	if hex := RGBToHex(c1.R, c1.G, c1.B); hex != expected {
		t.Errorf("expected %v got %v", expected, hex)
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
