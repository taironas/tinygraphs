package main

import (
	"flag"
	"github.com/taironas/route"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", "app", "file system path")

func main() {
	r := new(route.Router)

	r.HandleFunc("/black/?", blackHandler)
	r.HandleFunc("/green/?", greenHandler)
	r.HandleFunc("/blue/?", blueHandler)
	r.HandleFunc("/red/?", redHandler)
	r.HandleFunc("/grid/?", grid6X6Handler)
	r.HandleFunc("/gradient/?", gradientHandler)
	r.HandleFunc("/[1-8]/?", colorGridHandler)
	r.HandleFunc("/random/[1-8]/?", randomColorGridHandler)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
