package extract

import (
	"image/color"
	"net/http"
	"net/url"
	"testing"
)

func TestUserColors(t *testing.T) {
	tests := []struct {
		title  string
		url    string
		colors []color.RGBA
	}{
		{"test wrong input", "http://www.tg.c?colors=foo&colors=bar", []color.RGBA{}},
		{"test no input", "http://www.tg.c", []color.RGBA{}},
		{
			"test good input",
			"http://www.tg.c?colors=ffffff&colors=000000",
			[]color.RGBA{
				color.RGBA{255, 255, 255, 255},
				color.RGBA{0, 0, 0, 255},
			},
		},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		colors, err := UserColors(r)
		if err != nil {
			t.Log(err)
		}
		if len(colors) != len(test.colors) {
			t.Errorf("expected %d array got %d", len(test.colors), len(colors))
		}
		for i := 0; i < len(test.colors); i++ {
			if test.colors[i] != colors[i] {
				t.Errorf("expected %+v array got %+v", test.colors[i], colors[i])
			}
		}
	}
}
