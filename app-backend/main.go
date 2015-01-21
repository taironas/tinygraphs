package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/controllers/checkerboard"
	"github.com/taironas/tinygraphs/controllers/isogrids"
	"github.com/taironas/tinygraphs/controllers/squares"
)

var root = flag.String("root", "app", "file system path")

func main() {
	r := new(route.Router)

	r.HandleFunc("/checkerboard", checkerboard.Checkerboard)
	r.HandleFunc("/checkerboard/:colorId", checkerboard.Color)

	r.HandleFunc("/squares/?", squares.Random)
	r.HandleFunc("/squares/random/?", squares.Random)
	r.HandleFunc("/squares/random/:colorId", squares.RandomColor)
	r.HandleFunc("/squares/:key", squares.Square)         //cached
	r.HandleFunc("/squares/:colorId/:key", squares.Color) // cached

	r.HandleFunc("/isogrids/skeleton", isogrids.Skeleton)
	r.HandleFunc("/isogrids/:key", isogrids.Isogrids)
	r.HandleFunc("/isogrids/:colorId/:key", isogrids.Color)
	r.HandleFunc("/isogrids/diagonals", isogrids.Diagonals)
	r.HandleFunc("/isogrids/halfdiagonals", isogrids.HalfDiagonals)
	r.HandleFunc("/isogrids/gridbw", isogrids.GridBW)
	r.HandleFunc("/isogrids/grid2colors", isogrids.Grid2Colors)
	r.HandleFunc("/isogrids/random/:colorId", isogrids.RandomColor)
	r.HandleFunc("/isogrids/random", isogrids.Random)
	r.HandleFunc("/isogrids/random-mirror/:colorId", isogrids.RandomMirrorColor)
	r.HandleFunc("/isogrids/random-mirror", isogrids.RandomMirror)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
