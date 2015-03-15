package extract

import (
	"net/http"
	"net/url"
	"testing"
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
