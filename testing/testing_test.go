package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/taironas/route"
)

func TestGenerateHandlerFunc(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world")
	}

	r := new(route.Router)
	r.HandleFunc("/test", handler)

	test := GenerateHandlerFunc(t, handler)

	recorder := test("/test", "GET", map[string]string{}, r)
	if recorder.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
	}
}
