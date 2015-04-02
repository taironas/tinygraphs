package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taironas/route"
)

var GoodParams []map[string]string
var BadParams []map[string]string

func init() {
	GoodParams = []map[string]string{
		{},
		{"fmt": "jpeg"},
		{"fmt": "svg"},
		{"size": "200"},
		{"bg": "ff4008", "fg": "04d6f2"},
		{"theme": "frogideas"},
		{"w": "500", "h": "100"},
		{"xs": "500"},
		{"xt": "500"},
	}

	BadParams = []map[string]string{
		{"wrongfmt": "wrongParameter"},
		{"fmt": "wrongFormat"},
		{"size": "wrongType"},
		{"bg": "wrongType", "fg": "wrongType"},
		{"theme": "wrongTheme"},
		{"w": "wrongW", "h": "wrongH"},
		{"xs": "wrongXS"},
		{"xt": "wrongXT"},
	}
}

type HandlerFunc func(url, method string, params map[string]string, r *route.Router) *httptest.ResponseRecorder

func GenerateHandlerFunc(t *testing.T, handler func(http.ResponseWriter, *http.Request)) HandlerFunc {

	return func(url, method string, params map[string]string, r *route.Router) *httptest.ResponseRecorder {
		if req, err := http.NewRequest(method, url, nil); err != nil {
			t.Errorf("%v", err)
		} else {
			values := req.URL.Query()
			for k, v := range params {
				values.Add(k, v)
			}
			req.URL.RawQuery = values.Encode()
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)
			return recorder
		}
		return nil
	}
}
