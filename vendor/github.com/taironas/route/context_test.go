package route

import (
	"net/http"
	"testing"
)

func TestContext(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8080/", nil)

	var testContext context

	foo, _ := testContext.Get(r, "foo")
	if len(foo) > 0 {
		t.Fatal("Expected an empty string and got", foo)
	}

	params := make(map[string]string)
	params["foo"] = "johndoe"
	params["foo2"] = "42"

	testContext.set(r, params)

	if len(testContext.params[r]) != len(params) {
		t.Fatal("Params map should contained", len(params), "elements. context params length =", len(testContext.params[r]))
	}

	id, _ := testContext.Get(r, "foo2")
	if id != params["foo2"] {
		t.Fatal("Expected", params["foo2"], "and got", id)
	}

	foo3, err := testContext.Get(r, "foo3")
	if len(foo3) > 0 && err == nil {
		t.Fatal("Expected an empty string and got", foo3)
	}

	testContext.clear(r)
	if len(testContext.params) != 0 {
		t.Fatal("Params map should be empty")
	}
}
