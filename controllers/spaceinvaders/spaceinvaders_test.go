package spaceinvaders

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestSpaceInvaders(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/spaceinvaders/:key", SpaceInvaders)

	test := tgTesting.GenerateHandlerFunc(t, SpaceInvaders)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/spaceinvaders/somekey", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/spaceinvaders/somekey", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}

func TestSpaceInvadersCache(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/spaceinvaders/:key", SpaceInvaders)
	var etag string
	test := tgTesting.GenerateHandlerFunc(t, SpaceInvaders)

	if recorder := test("/spaceinvaders/somekey", "GET", nil, r); recorder != nil {
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		etag = recorder.Header().Get("Etag")
	}

	// test caching
	if req, err := http.NewRequest("GET", "/spaceinvaders/somekey", nil); err != nil {
		t.Errorf("%v", err)
	} else {
		req.Header.Set("If-None-Match", etag)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotModified {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusNotModified)
		}
	}
}
