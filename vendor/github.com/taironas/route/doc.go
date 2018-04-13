// Package route provides URL router allowing usage of regexp in URL paths.
//
//  package main
//
//  import (
//    "github.com/taironas/route"
//    "net/http"
//    "fmt"
//  )
//
//  func main() {
//    r := new(route.Router)
//    r.HandleFunc("/users", usersHandler)
//    r.HandleFunc("/users/:id", userHandler)
//    r.HandleFunc("/users/:id/friends/:username", friendHandler)
//
//    http.ListenAndServe(":8080", r)
//  }
//
//  func usersHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome to users handler!")
//  }
//
//  func userHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome to user handler, user id = %s!", route.Context.Get(r, "id"))
//  }
//
//  func friendHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome to friend handler, friend username = %s!", route.Context.Get(r, "username"))
//  }
package route
