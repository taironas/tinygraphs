package extract

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/taironas/tinygraphs/format"
)

func TestSize(t *testing.T) {
	tests := []struct {
		title string
		url   string
		size  int
	}{
		{"test wrong input", "http://www.tg.c?size=foo", 240},
		{"test no input", "http://www.tg.c", 240},
		{"test good input", "http://www.tg.c?size=10", 10},
		{"test lower limit", "http://www.tg.c?size=0", 240},
		{"test higher limit", "http://www.tg.c?size=1001", 240},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		s := Size(r)
		if s != test.size {
			t.Errorf("expected %d  got %d", test.size, s)
		}
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		title string
		url   string
		fmt   format.Format
	}{
		{"test wrong input", "http://www.tg.c?fmt=foo", format.JPEG},
		{"test no input", "http://www.tg.c", format.JPEG},
		{"test good input jpeg", "http://www.tg.c?fmt=jpeg", format.JPEG},
		{"test good input svg", "http://www.tg.c?fmt=svg", format.SVG},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		f := Format(r)
		if f != test.fmt {
			t.Errorf("expected %d got %d", test.fmt, f)
		}
	}
}

func TestTheme(t *testing.T) {
	tests := []struct {
		title string
		url   string
		theme string
	}{
		{"test wrong input", "http://www.tg.c?fmt=", "base"},
		{"test no input", "http://www.tg.c", "base"},
		{"test good input", "http://www.tg.c?theme=hello", "hello"},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		th := Theme(r)
		if th != test.theme {
			t.Errorf("expected %d got %d", test.theme, th)
		}
	}
}

func TestHexalines(t *testing.T) {
	tests := []struct {
		title     string
		url       string
		hexalines int
	}{
		{"test wrong input", "http://www.tg.c?hexalines=h", 6},
		{"test no input", "http://www.tg.c", 6},
		{"test good input", "http://www.tg.c?hexalines=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		x := Hexalines(r)
		if x != test.hexalines {
			t.Errorf("expected %d got %d", test.hexalines, x)
		}
	}
}

func TestLines(t *testing.T) {
	tests := []struct {
		title string
		url   string
		lines int
	}{
		{"test wrong input", "http://www.tg.c?lines=h", 6},
		{"test no input", "http://www.tg.c", 6},
		{"test good input", "http://www.tg.c?lines=4", 4},
		{"test input limit", "http://www.tg.c?lines=3", 6},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		l := Lines(r)
		if l != test.lines {
			t.Errorf("expected %d got %d", test.lines, l)
		}
	}
}

func TestWidth(t *testing.T) {
	tests := []struct {
		title string
		url   string
		width int
	}{
		{"test wrong input", "http://www.tg.c?w=", 720},
		{"test no input", "http://www.tg.c", 720},
		{"test good input", "http://www.tg.c?w=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		w := Width(r)
		if w != test.width {
			t.Errorf("expected %d got %d", test.width, w)
		}
	}
}

func TestWidthOrDefault(t *testing.T) {
	tests := []struct {
		title string
		url   string
		width int
	}{
		{"test wrong input", "http://www.tg.c?w=", 1},
		{"test no input", "http://www.tg.c", 1},
		{"test good input", "http://www.tg.c?w=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		w := WidthOrDefault(r, 1)
		if w != test.width {
			t.Errorf("expected %d got %d", test.width, w)
		}
	}
}

func TestHeight(t *testing.T) {
	tests := []struct {
		title  string
		url    string
		height int
	}{
		{"test wrong input", "http://www.tg.c?h=", 300},
		{"test no input", "http://www.tg.c", 300},
		{"test good input", "http://www.tg.c?h=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		h := Height(r)
		if h != test.height {
			t.Errorf("expected %d got %d", test.height, h)
		}
	}
}

func TestHeightOrDefault(t *testing.T) {
	tests := []struct {
		title  string
		url    string
		height int
	}{
		{"test wrong input", "http://www.tg.c?h=", 1},
		{"test no input", "http://www.tg.c", 1},
		{"test good input", "http://www.tg.c?h=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		h := HeightOrDefault(r, 1)
		if h != test.height {
			t.Errorf("expected %d got %d", test.height, h)
		}
	}
}

func TestXSquares(t *testing.T) {
	tests := []struct {
		title string
		url   string
		xs    int
	}{
		{"test wrong input", "http://www.tg.c?xs=hello", 50},
		{"test no input", "http://www.tg.c", 50},
		{"test good input", "http://www.tg.c?xs=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		xs := XSquares(r)
		if xs != test.xs {
			t.Errorf("expected %d got %d", test.xs, xs)
		}
	}
}

func TestXTriangles(t *testing.T) {
	tests := []struct {
		title string
		url   string
		xt    int
	}{
		{"test wrong input", "http://www.tg.c?xt=hello", 50},
		{"test no input", "http://www.tg.c", 50},
		{"test good input", "http://www.tg.c?xt=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		xt := XTriangles(r)
		if xt != test.xt {
			t.Errorf("expected %d got %d", test.xt, xt)
		}
	}
}

func TestGX1OrDefault(t *testing.T) {
	tests := []struct {
		title string
		url   string
		gx1   uint8
	}{
		{"test wrong input", "http://www.tg.c?gx1=hello", 1},
		{"test no input", "http://www.tg.c", 1},
		{"test good input", "http://www.tg.c?gx1=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		gx1 := GX1OrDefault(r, uint8(1))
		if gx1 != test.gx1 {
			t.Errorf("expected %d got %d", test.gx1, gx1)
		}
	}
}

func TestGX2OrDefault(t *testing.T) {
	tests := []struct {
		title string
		url   string
		gx2   uint8
	}{
		{"test wrong input", "http://www.tg.c?gx2=hello", 1},
		{"test no input", "http://www.tg.c", 1},
		{"test good input", "http://www.tg.c?gx2=4", 4},
	}

	for _, test := range tests {
		t.Log(test.title)
		r := &http.Request{Method: "GET"}
		r.URL, _ = url.Parse(test.url)
		gx2 := GX2OrDefault(r, uint8(1))
		if gx2 != test.gx2 {
			t.Errorf("expected %d got %d", test.gx2, gx2)
		}
	}
}
