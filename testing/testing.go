package testing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var GoodParams []string
var BadParams []string

func init() {
	GoodParams = []string{`{}`,
		`{"fmt":"jpeg"}`,
		`{"fmt":"svg"}`,
		`{"size":"200"}`,
		`{"bg":"ff4008","fg":"04d6f2"}`,
		`{"theme":"frogideas"}`,
	}

	BadParams = []string{
		`{"wrongfmt":"wrongParameter"}`,
		`{"fmt":"wrongFormat"}`,
		`{"size":"wrongType"}`,
		`{"bg":"wrongType","fg":"wrongType"}`,
		`{"theme":"wrongTheme"}`,
	}
}

type HandlerFunc func(method, params string) *httptest.ResponseRecorder

func GenerateHandlerFunc(t *testing.T, handler func(http.ResponseWriter, *http.Request)) HandlerFunc {

	return func(method, params string) *httptest.ResponseRecorder {

		if r, err := http.NewRequest(method, "", strings.NewReader(params)); err != nil {
			t.Errorf("%v", err)
		} else {
			r.Header.Set("Content-Type", "application/json")
			r.Body.Close()
			recorder := httptest.NewRecorder()
			handler(recorder, r)
			return recorder
		}
		return nil
	}
}
