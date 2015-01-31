package isogrids

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestHexa(t *testing.T) {

	r := new(route.Router)
	r.HandleFunc("/isogrids/labs/hexa", Hexa)
	r.HandleFunc("/isogrids/labs/hexa/:key", Hexa)

	test := tgTesting.GenerateHandlerFunc(t, Hexa)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/isogrids/labs/hexa", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		recorder = test("/isogrids/labs/hexa/somekey", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}

	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/isogrids/labs/hexa", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		recorder = test("/isogrids/labs/hexa/somekey", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}

	}
}
