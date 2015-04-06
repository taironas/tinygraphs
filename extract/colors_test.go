package extract

import (
	"image/color"
	"net/http"
	"net/url"
	"testing"

	tgColors "github.com/taironas/tinygraphs/colors"
)

func TestColors(t *testing.T) {
	t.Parallel()
	colorMap := tgColors.MapOfColorThemes()

	tests := []struct {
		title  string
		url    string
		colors []color.RGBA
	}{
		{
			"test wrong input",
			"http://www.tg.c?colors=foo&colors=bar",
			colorMap["base"],
		},
		{
			"test no input",
			"http://www.tg.c",
			colorMap["base"],
		},
		{
			"test good input",
			"http://www.tg.c?colors=ffffff&colors=000000",
			[]color.RGBA{
				{255, 255, 255, 255},
				{0, 0, 0, 255},
			},
		},
		{
			"test good input with theme",
			"http://www.tg.c?theme=frogideas",
			colorMap["frogideas"][0:2],
		},
		{
			"test good input with theme and num color",
			"http://www.tg.c?theme=frogideas&numcolors=4",
			colorMap["frogideas"],
		},
		{
			"test good input with theme and order array",
			"http://www.tg.c?theme=frogideas&numcolors=4&order=3&order=2&order=1&order=0",
			[]color.RGBA{
				colorMap["frogideas"][3],
				colorMap["frogideas"][2],
				colorMap["frogideas"][1],
				colorMap["frogideas"][0],
			},
		},

		{
			"test bad theme",
			"http://www.tg.c?theme=bad",
			colorMap["base"],
		},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		colors := Colors(r)
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

func TestUserColors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		url    string
		colors []color.RGBA
	}{
		{
			"test wrong input",
			"http://www.tg.c?colors=foo&colors=bar",
			[]color.RGBA{},
		},
		{
			"test no input",
			"http://www.tg.c",
			[]color.RGBA{},
		},
		{
			"test good input",
			"http://www.tg.c?colors=ffffff&colors=000000",
			[]color.RGBA{
				{255, 255, 255, 255},
				{0, 0, 0, 255},
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

func TestGColors(t *testing.T) {
	t.Parallel()
	colorMap := tgColors.MapOfColorThemes()

	tests := []struct {
		title   string
		url     string
		gColors []color.RGBA
	}{
		{
			"test wrong input",
			"http://www.tg.c?colors=foo&colors=bar",
			colorMap["base"],
		},
		{
			"test no input",
			"http://www.tg.c",
			colorMap["base"],
		},
		{
			"test good input",
			"http://www.tg.c?colors=aaaaaa&colors=bbbbbb",
			[]color.RGBA{
				{170, 170, 170, 255},
				{187, 187, 187, 255},
			},
		},
		{
			"test good input",
			"http://www.tg.c?colors=ffffff&colors=000000&colors=000000",
			[]color.RGBA{
				{255, 255, 255, 255},
				{0, 0, 0, 255},
				{0, 0, 0, 255},
			}[1:3],
		},
		{
			"test with theme",
			"http://www.tg.c?theme=frogideas",
			colorMap["frogideas"][1:3],
		},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		gColors := GColors(r)
		if len(gColors) != len(test.gColors) {
			t.Errorf("expected %d array got %d", len(test.gColors), len(gColors))
		}
		for i := 0; i < len(test.gColors); i++ {
			if test.gColors[i] != gColors[i] {
				t.Errorf("expected %v array got %v", test.gColors[i], gColors[i])
			}
		}
	}
}

func TestBackground(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title string
		url   string
		bg    color.RGBA
	}{
		{
			"test wrong input",
			"http://www.tg.c?bg=foo",
			color.RGBA{},
		},
		{
			"test no input",
			"http://www.tg.c",
			color.RGBA{},
		},
		{
			"test good input",
			"http://www.tg.c?bg=aaaaaa",
			color.RGBA{170, 170, 170, 255},
		},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		bg, _ := Background(r)
		if bg != test.bg {
			t.Errorf("expected %+v got %+v", test.bg, bg)
		}
	}
}

func TestForeground(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title string
		url   string
		fg    color.RGBA
	}{
		{
			"test wrong input",
			"http://www.tg.c?fg=foo",
			color.RGBA{},
		},
		{
			"test no input",
			"http://www.tg.c",
			color.RGBA{},
		},
		{
			"test good input",
			"http://www.tg.c?fg=aaaaaa",
			color.RGBA{170, 170, 170, 255},
		},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		fg, _ := Foreground(r)
		if fg != test.fg {
			t.Errorf("expected %+v got %+v", test.fg, fg)
		}
	}
}

func TestExtraColors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title string
		url   string
		bg    color.RGBA
		fg    color.RGBA
	}{
		{
			"test wrong input",
			"http://www.tg.c?fg=foo&bg=bar",
			tgColors.White(),
			tgColors.Black(),
		},
		{
			"test no input",
			"http://www.tg.c",
			tgColors.White(),
			tgColors.Black(),
		},
		{
			"test good input",
			"http://www.tg.c?fg=aaaaaa&bg=bbbbbb",
			color.RGBA{187, 187, 187, 255},
			color.RGBA{170, 170, 170, 255},
		},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		bg, fg := ExtraColors(r)
		if fg != test.fg || bg != test.bg {
			t.Errorf("expected %+v got %+v", test.fg, fg)
			t.Errorf("expected %+v got %+v", test.bg, bg)
		}
	}
}

func TestReOrder(t *testing.T) {
	t.Parallel()
	colors := []color.RGBA{
		{255, 255, 255, 255},
		{0, 0, 0, 255},
		{0, 0, 0, 255},
		{255, 255, 255, 255},
	}
	expected := []color.RGBA{
		{0, 0, 0, 255},
		{0, 0, 0, 255},
		{255, 255, 255, 255},
		{255, 255, 255, 255},
	}
	ReOrder([]int{1, 2, 0, 3}, &colors)
	if !areArrayOfColorsEqual(colors, expected) {
		t.Errorf("expected %v got %v", expected, colors)
	}

	// test wrong order array does not change colors array
	colors = []color.RGBA{
		{255, 255, 255, 255},
		{0, 0, 0, 255},
		{0, 0, 0, 255},
		{255, 255, 255, 255},
	}
	expected = colors
	ReOrder([]int{1, 2, 0, 4}, &colors)
	if !areArrayOfColorsEqual(colors, expected) {
		t.Errorf("expected %v got %v", expected, colors)
	}

}

func TestSwap(t *testing.T) {
	t.Parallel()
	colors := []color.RGBA{
		{255, 255, 255, 255},
		{0, 0, 0, 255},
	}
	expected := []color.RGBA{
		{0, 0, 0, 255},
		{255, 255, 255, 255},
	}
	swap(&colors)
	if !areArrayOfColorsEqual(colors, expected) {
		t.Errorf("expected %v got %v", expected, colors)
	}
}

func areArrayOfColorsEqual(a, b []color.RGBA) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
