package isogrids

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestIsogrids(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/isogrids/:key", Isogrids)

	test := tgTesting.GenerateHandlerFunc(t, Isogrids)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/isogrids/test", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/isogrids/test", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}

func BenchmarkIsogrids(b *testing.B) {
	r := new(route.Router)
	r.HandleFunc("/isogrids/:key", Isogrids)

	request := req(b, "GET /isogrids/test HTTP/1.0\r\n\r\n")
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
