package isogrids

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestLabs(t *testing.T) {

	r := new(route.Router)

	type Labs struct {
		url     string
		handler func(http.ResponseWriter, *http.Request)
	}

	labs := [...]Labs{
		{"/isogrids/labs/skeleton", Skeleton},
	}
	for _, lab := range labs {
		r.HandleFunc(lab.url, lab.handler)

		test := tgTesting.GenerateHandlerFunc(t, lab.handler)
		for _, p := range tgTesting.GoodParams {
			recorder := test(lab.url, "GET", p, r)
			if recorder.Code != http.StatusOK {
				t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
			}

		}

		for _, p := range tgTesting.BadParams {
			recorder := test(lab.url, "GET", p, r)
			if recorder.Code != http.StatusOK {
				t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
			}
		}

	}
}
