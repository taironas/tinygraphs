package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", "app", "file system path")

func main() {
	http.HandleFunc("/black/", blackHandler)
	http.HandleFunc("/green/", greenHandler)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func blackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "black handler!")
}

func greenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "green handler!")
}
