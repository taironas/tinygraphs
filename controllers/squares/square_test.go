package squares

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/taironas/route"

	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestSquares(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/squares/:key", Square)

	test := tgTesting.GenerateHandlerFunc(t, Square)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/squares/test", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/squares/test", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	// test caching:
	recorder := test("/squares/cache", "GET", nil, r)
	if recorder.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
	}

	if req, err := http.NewRequest("GET", "/squares/cache", nil); err != nil {
		t.Errorf("%v", err)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("If-None-Match", recorder.Header().Get("Etag"))
		recorder2 := httptest.NewRecorder()
		r.ServeHTTP(recorder2, req)
		if recorder2.Code != http.StatusNotModified {
			t.Errorf("returned %v. Expected %v.", recorder2.Code, http.StatusNotModified)
		}
	}
}

func BenchmarkSquares(b *testing.B) {
	r := new(route.Router)
	r.HandleFunc("/squares/:key", Square)

	request := req(b, "GET /squares/test HTTP/1.0\r\n\r\n")
	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)
	}
}

func req(tb testing.TB, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		tb.Fatal(err)
	}
	return req
}
