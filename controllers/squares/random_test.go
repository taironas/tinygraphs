package squares

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestRandom(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/squares/random", Random)

	test := tgTesting.GenerateHandlerFunc(t, Random)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/squares/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/squares/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}
