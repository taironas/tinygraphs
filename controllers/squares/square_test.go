package squares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taironas/route"

	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestSquares(t *testing.T) {

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
