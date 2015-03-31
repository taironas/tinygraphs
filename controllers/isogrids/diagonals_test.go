package isogrids

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestDiagonals(t *testing.T) {
	t.Parallel()
	r := new(route.Router)

	type Diagonal struct {
		url     string
		handler func(http.ResponseWriter, *http.Request)
	}

	diagonals := [...]Diagonal{
		{"/isogrids/labs/diagonals", Diagonals},
		{"/isogrids/labs/halfdiagonals", HalfDiagonals},
	}
	for _, d := range diagonals {
		r.HandleFunc(d.url, d.handler)
		test := tgTesting.GenerateHandlerFunc(t, d.handler)

		for _, p := range tgTesting.GoodParams {
			recorder := test(d.url, "GET", p, r)
			if recorder.Code != http.StatusOK {
				t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
			}
		}

		for _, p := range tgTesting.BadParams {
			recorder := test(d.url, "GET", p, r)
			if recorder.Code != http.StatusOK {
				t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
			}
		}
	}
}
