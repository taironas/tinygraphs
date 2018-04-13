package route

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var endpoints = []string{
	"/test/handler/1",
	"/test/handler/2",
}

var rootTestingPath = "temp_TestServeStaticResources"

func TestRouter(t *testing.T) {
	r := new(Router)
	r.HandleFunc(endpoints[0], handler1)
	r.HandleFunc(endpoints[1], handler2)

	server := httptest.NewServer(r)
	defer server.Close()

	var urls = []string{
		"/test/handler/1",
		"/test/handler/2",
		"/test/handler/1/",
		"/test/handler/2/",
	}

	for _, u := range urls {
		res, err := http.Get(server.URL + u)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatal(res)
		}
	}
}

func TestNotFoundRoute(t *testing.T) {
	r := new(Router)
	r.HandleFunc(endpoints[0], handler1)
	r.HandleFunc(endpoints[1], handler2)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler/3")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}
}

func TestFoundRouteWithVariables(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/:id/hello/", handlerHello)
	r.HandleFunc("/test/handler/:id/hello/:foo", handlerHello2)

	urls := []string{
		"/test/handler/1/hello",
		"/test/handler/2/hello/",
		"/test/handler/2123/hello/johndoe",
		"/test/handler/2123/hello/johndoe/",
	}

	server := httptest.NewServer(r)
	defer server.Close()

	for _, u := range urls {
		res, err := http.Get(server.URL + u)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Fatal(res)
		}
	}

	res, _ := http.Get(server.URL + urls[1])
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	expectedBody := "2"
	gotBody := string(body)
	if gotBody != expectedBody {
		t.Fatal("Expected", expectedBody, "and got", gotBody)
	}

	res, _ = http.Get(server.URL + urls[2])
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	expectedBody = "2123,johndoe"
	gotBody = string(body)
	if gotBody != expectedBody {
		t.Fatal("Expected", expectedBody, "and got", gotBody)
	}
}

func TestNotFoundRouteWithVariables(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/:id/hello", handlerHello)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler///hello")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}
}

func TestFoundRouteWithSpecialCharacter(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/:user_id/hello/:f-o-o", handlerHelloWithSpecialCharacter)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler/42/hello/johndoe")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal(res)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	expectedBody := "42,johndoe"
	gotBody := string(body)
	if gotBody != expectedBody {
		t.Fatal("Expected", expectedBody, "and got", gotBody)
	}
}

func TestNotFoundRouteWithSpecialCharacter(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/:user_id/hello/:f+o+o", handlerHelloWithSpecialCharacter)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler/42/hello/johndoe")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}
}

func TestServeStaticResources(t *testing.T) {

	createTestingData(rootTestingPath)

	r := new(Router)
	r.AddStaticResource(&rootTestingPath)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal(res)
	}

	cleanTestingData(rootTestingPath)
}

func TestServeTwoLevelStaticResources(t *testing.T) {
	cssTestingPath := rootTestingPath + "/css"

	createTestingData(rootTestingPath)

	r := new(Router)
	r.AddStaticResource(&cssTestingPath)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal(res)
	}

	cleanTestingData(rootTestingPath)
}

func TestServeNonExistingStaticResources(t *testing.T) {
	jsTestingPath := rootTestingPath + "/js"

	createTestingData(rootTestingPath)

	r := new(Router)
	r.AddStaticResource(&jsTestingPath)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}

	cleanTestingData(rootTestingPath)
}

func TestMatch(t *testing.T) {
	params := make(map[string]string)
	routeWithOneVariable := route{"/test/handler/:id/hello", params, nil}
	routeWithMultipleVariables := route{"/test/handler/:id/hello/:username", params, nil}

	matchingPattern := "/test/handler/50/hello"
	if !routeWithOneVariable.match(matchingPattern) {
		t.Fatal("route should match the pattern: pattern = " + routeWithOneVariable.pattern + ", path = " + matchingPattern)
	}

	if routeWithOneVariable.params["id"] != "50" {
		t.Fatal("Value for 'user_id' is not the expected one: expected = 50, stored = " + routeWithOneVariable.params["id"])
	}

	nonMatchingPattern := "/test/handler//hello"
	if !routeWithOneVariable.match(nonMatchingPattern) {
		t.Fatal("route should not match the pattern: pattern = " + routeWithOneVariable.pattern + ", path = " + nonMatchingPattern)
	}

	matchingPattern = "/test/handler/johndoe/hello"
	if !routeWithOneVariable.match(matchingPattern) {
		t.Fatal("route should match the pattern: pattern = " + routeWithOneVariable.pattern + ", path = " + matchingPattern)
	}

	if routeWithOneVariable.params["id"] != "johndoe" {
		t.Fatal("Value for 'user_id' is not the expected one: expected = johndoe, stored = " + routeWithOneVariable.params["id"])
	}

	matchingPattern = "/test/handler/50/hello/johndoe"
	if !routeWithMultipleVariables.match(matchingPattern) {
		t.Fatal("route should match the pattern: pattern = " + routeWithMultipleVariables.pattern + ", path = " + matchingPattern)
	}

	if routeWithMultipleVariables.params["id"] != "50" && routeWithMultipleVariables.params["username"] != "johndoe" {
		t.Fatal("Values stored in map are not the expected ones: expected = [50, john doe], stored = [" + routeWithMultipleVariables.params["id"] + ", " + routeWithMultipleVariables.params["username"] + "]")
	}

	nonMatchingPattern = "/test/handler/:id/hello//"
	if !routeWithMultipleVariables.match(nonMatchingPattern) {
		t.Fatal("route should not match the pattern: pattern = " + routeWithMultipleVariables.pattern + ", path = " + nonMatchingPattern)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler1 has been reached!")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler2 has been reached!")
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	id, _ := Context.Get(r, "id")
	fmt.Fprintf(w, id)
}

func handlerHello2(w http.ResponseWriter, r *http.Request) {
	id, _ := Context.Get(r, "id")
	foo, _ := Context.Get(r, "foo")
	fmt.Fprintf(w, id+","+foo)
}

func handlerHelloWithSpecialCharacter(w http.ResponseWriter, r *http.Request) {
	userID, _ := Context.Get(r, "user_id")
	foo, _ := Context.Get(r, "f-o-o")
	fmt.Fprintf(w, userID+","+foo)
}

func createTestingData(rootTestPath string) {
	src, err := os.Stat(rootTestPath)
	if err != nil || !src.IsDir() {
		os.Mkdir(rootTestPath, 0777)
	}

	src, err = os.Stat(rootTestPath + "/index.html")
	if err != nil || src.IsDir() {
		os.Create(rootTestPath + "/index.html")
	}

	cssTestPath := rootTestPath + "/css"

	src, err = os.Stat(cssTestPath)
	if err != nil || !src.IsDir() {
		os.Mkdir(cssTestPath, 0777)
	}

	src, err = os.Stat(cssTestPath + "/main.css")
	if err != nil || src.IsDir() {
		os.Create(cssTestPath + "/main.css")
	}
}

func cleanTestingData(nameTest string) {
	src, err := os.Stat(nameTest)
	if err == nil && src.IsDir() {
		os.RemoveAll(nameTest)
	}
}
