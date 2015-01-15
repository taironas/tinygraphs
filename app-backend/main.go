package main

import (
	"flag"
	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/controllers/grid"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", "app", "file system path")

func main() {
	r := new(route.Router)

	r.HandleFunc("/grid/?", grid.H6X6)
	r.HandleFunc("/grid/[0-8]/?", grid.Color)

	r.HandleFunc("/squares/?", grid.Random)
	r.HandleFunc("/squares/random/?", grid.Random)
	r.HandleFunc("/squares/random/[0-8]/?", grid.RandomColor)
	r.HandleFunc("/squares/[a-zA-Z0-9]+/?", grid.Square)            //cached
	r.HandleFunc("/squares/[0-8]/[a-zA-Z0-9]+/?", grid.SquareColor) // cached

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
