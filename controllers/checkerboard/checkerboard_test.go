package checkerboard

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestCheckerboard(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/checkerboard", Checkerboard)

	test := tgTesting.GenerateHandlerFunc(t, Checkerboard)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/checkerboard", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/checkerboard", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

}
