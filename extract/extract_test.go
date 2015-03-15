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
