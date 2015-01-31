package checkerboard

import (
	"net/http"
	"testing"

	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestCheckerboard(t *testing.T) {

	test := tgTesting.GenerateHandlerFunc(t, Checkerboard)
	for _, p := range tgTesting.GoodParams {
		recorder := test("GET", p)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusBadRequest)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("GET", p)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusBadRequest)
		}
	}

}
