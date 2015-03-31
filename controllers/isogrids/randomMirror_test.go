package isogrids

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestRandomMirror(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/isogrids/random-mirror", RandomMirror)

	test := tgTesting.GenerateHandlerFunc(t, RandomMirror)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/isogrids/random-mirror", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/isogrids/random-mirror", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}
