package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// parse permalink id from URL  and return it
func PermalinkID(r *http.Request, level int64) (int64, error) {
	url := strings.Replace(r.URL.String(), "http://", "", 1)
	path := strings.Split(url, "/")
	var strID string
	strID = path[level]
	intID, err := strconv.ParseInt(strID, 0, 64)
	if err != nil {
		// only try to extract id if were are unable to exracted using the level.
		if strings.Contains(r.URL.String(), "?") {
			strPath := path[level]
			strID = strPath[0:strings.Index(strPath, "?")]
		} else {
			strID = path[level]
		}
		intID, err = strconv.ParseInt(strID, 0, 64)
		if err != nil {
			log.Printf("error when calling PermalinkID with %v.Error: %v", path[level], err)
		}
	}
	return intID, err
}
