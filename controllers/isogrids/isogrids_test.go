package isogrids

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestIsogrids(t *testing.T) {

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
