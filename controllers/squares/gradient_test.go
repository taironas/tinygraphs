package squares

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestBannerGradient(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/squares/banner/gradient", BannerGradient)

	test := tgTesting.GenerateHandlerFunc(t, BannerGradient)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/squares/banner/gradient", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/squares/banner/gradient", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}
