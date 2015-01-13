package main

import (
	"flag"
	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/controllers/colors"
	"github.com/taironas/tinygraphs/controllers/gradient"
	"github.com/taironas/tinygraphs/controllers/grid"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", "app", "file system path")

func main() {
	r := new(route.Router)

	r.HandleFunc("/black/?", colors.Black)
	r.HandleFunc("/green/?", colors.Green)
	r.HandleFunc("/blue/?", colors.Blue)
	r.HandleFunc("/red/?", colors.Red)

	r.HandleFunc("/grid/?", grid.H6X6)
	r.HandleFunc("/grid/[0-8]/?", grid.Color)
	r.HandleFunc("/grid/square/[0-8]/[a-zA-Z0-9]+/?", grid.SquareColor)
	r.HandleFunc("/grid/square/[a-zA-Z0-9]+/?", grid.Square)

	r.HandleFunc("/grid/random/?", grid.Random)
	r.HandleFunc("/grid/random/[0-8]/?", grid.RandomColor)

	r.HandleFunc("/grid/random/symetric/?", grid.RandomSymetricX)
	r.HandleFunc("/grid/random/symetric/x/?", grid.RandomSymetricX)
	r.HandleFunc("/grid/random/symetric/x/[0-8]?", grid.RandomSymetricXColor)
	r.HandleFunc("/grid/random/symetric/y/?", grid.RandomSymetricY)
	r.HandleFunc("/grid/random/symetric/y/[0-8]?", grid.RandomSymetricYColor)

	r.HandleFunc("/gradient/?", gradient.Gradient)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
