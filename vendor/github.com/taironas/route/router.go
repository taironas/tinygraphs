package route

import (
	"net/http"
	"regexp"
	"strings"
)

// inspired by the following sources with some small changes:
// http://stackoverflow.com/questions/6564558/wildcards-in-the-pattern-for-http-handlefunc
// https://github.com/raymi/quickerreference
type route struct {
	pattern string
	params  map[string]string
	handler http.Handler
}

// Router serves HTTP requests for added routes and static resources
type Router struct {
	routes          []*route  // array of routes with a tuple (pattern, handler)
	staticResources []*string // array of static resources
}

// Context holds URL parameters indexed by HTTP request
var Context context

// Handle registers the handler for the given pattern in the router.
func (r *Router) Handle(pattern string, handler http.Handler) {
	params := make(map[string]string)
	r.routes = append(r.routes, &route{pattern, params, handler})
}

// HandleFunc registers the handler function for the given pattern in the router.
func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.Handle(pattern, clearHandler(http.HandlerFunc(handler)))
}

// ServeHTTP looks for a matching route among the routes. Returns 404 if no match is found.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.match(req.URL.Path) {
			Context.set(req, route.params)
			route.handler.ServeHTTP(w, req)
			return
		}
	}

	// route not found. check if it is a static ressource.
	for _, sr := range r.staticResources {
		dir := http.Dir(*sr)
		if _, err := dir.Open(req.URL.Path); err == nil {
			// Could open file, set static resource and call ServeHTTP again.
			r.Handle(req.URL.Path, http.FileServer(dir))
			r.ServeHTTP(w, req)
			return
		}
	}

	// no pattern matched; send 404 response
	http.NotFound(w, req)
}

// AddStaticResource adds a resource value to an array of static resources.
// Use this if you want to serve a static directory and it's sub directories.
func (r *Router) AddStaticResource(resource *string) {
	r.staticResources = append(r.staticResources, resource)
}

// Determines if the path matches the pattern of the route.
// Fill the params map with variables found in the path.
func (r *route) match(path string) bool {
	splitPattern := strings.Split(strings.TrimSuffix(r.pattern, "/"), "/")
	splitPath := strings.Split(strings.TrimSuffix(path, "/"), "/")

	if len(splitPattern) != len(splitPath) {
		return false
	}

	for idx, val := range splitPattern {
		if val != splitPath[idx] {
			if !strings.HasPrefix(val, ":") {
				return false
			}

			variable := strings.TrimPrefix(val, ":")
			if matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", variable); !matched {
				return false
			}

			r.params[variable] = splitPath[idx]
		}
	}

	return true
}

// Handler clears the context for a given request.
func clearHandler(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer Context.clear(r)
		f(w, r)
	}
}
