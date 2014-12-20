package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", "app", "file system path")

func main() {
	http.HandleFunc("/black/", blackHandler)
	http.HandleFunc("/green/", greenHandler)
	http.HandleFunc("/blue/", blueHandler)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
