package spaceinvaders

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestSpaceInvaders(t *testing.T) {

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
