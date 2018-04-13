package route

import (
	"errors"
	"net/http"
	"sync"
)

// inspired by Brad Fitzpatrick's idea:
// https://groups.google.com/forum/#!msg/golang-nuts/teSBtPvv1GQ/U12qA9N51uIJ
type context struct {
	mutex  sync.Mutex
	params map[*http.Request]map[string]string // URL parameters.
}

// set stores a map of URL paramters for a given request.
func (c *context) set(req *http.Request, m map[string]string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.params == nil {
		c.params = make(map[*http.Request]map[string]string)
	}
	c.params[req] = m
}

// Get returns an URL parameter value for a given key for a given request.
func (c *context) Get(req *http.Request, key string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.params == nil {
		return "", errors.New("Parameters map has not been initialized")
	}

	val, ok := c.params[req][key]

	if !ok {
		return val, errors.New(key + " Key does not exist in the parameters map")
	}
	return val, nil
}

//  clear removes all the key/value pairs for a given request.
func (c *context) clear(req *http.Request) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.params, req)
}
