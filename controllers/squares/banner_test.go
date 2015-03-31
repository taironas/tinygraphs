package squares

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestBannerRandom(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/squares/banner/random", BannerRandom)

	test := tgTesting.GenerateHandlerFunc(t, BannerRandom)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/squares/banner/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/squares/banner/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}

func TestBannerRandomGradient(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/squares/banner/random", BannerRandomGradient)

	test := tgTesting.GenerateHandlerFunc(t, BannerRandomGradient)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/squares/banner/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/squares/banner/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}
