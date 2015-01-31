package squares

import (
	"net/http"
	"testing"

	"github.com/taironas/route"

	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestSquares(t *testing.T) {

	r := new(route.Router)
	r.HandleFunc("/squares/:key", Square)

	test := tgTesting.GenerateHandlerFunc(t, Square)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/squares/test", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusBadRequest)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/squares/test", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusBadRequest)
		}
	}

}
