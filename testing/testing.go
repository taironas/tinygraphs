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
		map[string]string{},
		map[string]string{"fmt": "jpeg"},
		map[string]string{"fmt": "svg"},
		map[string]string{"size": "200"},
		map[string]string{"bg": "ff4008", "fg": "04d6f2"},
		map[string]string{"theme": "frogideas"},
		map[string]string{"w": "500", "h": "100"},
		map[string]string{"xs": "500"},
		map[string]string{"xt": "500"},
	}

	BadParams = []map[string]string{
		map[string]string{"wrongfmt": "wrongParameter"},
		map[string]string{"fmt": "wrongFormat"},
		map[string]string{"size": "wrongType"},
		map[string]string{"bg": "wrongType", "fg": "wrongType"},
		map[string]string{"theme": "wrongTheme"},
		map[string]string{"w": "wrongW", "h": "wrongH"},
		map[string]string{"xs": "wrongXS"},
		map[string]string{"xt": "wrongXT"},
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
