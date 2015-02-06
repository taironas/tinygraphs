package testing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/taironas/route"
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
		`{"w":"500", "h":"100"}`,
		`{"xs":"500"}`,
		`{"xt":"500"}`,
	}

	BadParams = []string{
		`{"wrongfmt":"wrongParameter"}`,
		`{"fmt":"wrongFormat"}`,
		`{"size":"wrongType"}`,
		`{"bg":"wrongType","fg":"wrongType"}`,
		`{"theme":"wrongTheme"}`,
		`{"w":"wrongW", "h":"wrongH"}`,
		`{"xs":"wrongXS"}`,
		`{"xt":"wrongXT"}`,
	}
}

type HandlerFunc func(url, method, params string, r *route.Router) *httptest.ResponseRecorder

func GenerateHandlerFunc(t *testing.T, handler func(http.ResponseWriter, *http.Request)) HandlerFunc {

	return func(url, method, params string, r *route.Router) *httptest.ResponseRecorder {

		if req, err := http.NewRequest(method, url, strings.NewReader(params)); err != nil {
			t.Errorf("%v", err)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Body.Close()
			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)
			return recorder
		}
		return nil
	}
}
