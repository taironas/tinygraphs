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

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

func main() {
	r := new(route.Router)

	r.HandleFunc("/checkerboard", checkerboard.Checkerboard)

	r.HandleFunc("/squares", squares.Random)
	r.HandleFunc("/squares/random", squares.Random)
	r.HandleFunc("/squares/:key", squares.Square) //cached

	r.HandleFunc("/isogrids/random", isogrids.Random)
	r.HandleFunc("/isogrids/random-mirror", isogrids.RandomMirror)

	r.HandleFunc("/isogrids/:key", isogrids.Isogrids)

	r.HandleFunc("/isogrids/labs/skeleton", isogrids.Skeleton)
	r.HandleFunc("/isogrids/labs/diagonals", isogrids.Diagonals)
	r.HandleFunc("/isogrids/labs/halfdiagonals", isogrids.HalfDiagonals)
	r.HandleFunc("/isogrids/labs/gridbw", isogrids.GridBW)
	r.HandleFunc("/isogrids/labs/grid2colors", isogrids.Grid2Colors)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
