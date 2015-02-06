package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/controllers/checkerboard"
	"github.com/taironas/tinygraphs/controllers/isogrids"
	"github.com/taironas/tinygraphs/controllers/spaceinvaders"
	"github.com/taironas/tinygraphs/controllers/squares"
)

var root = flag.String("root", "app", "file system path")

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

func main() {
	r := new(route.Router)

	r.HandleFunc("/squares", squares.Random)
	r.HandleFunc("/squares/banner/random", squares.BannerRandom)
	r.HandleFunc("/squares/banner/random/gradient", squares.BannerRandomGradient)
	r.HandleFunc("/squares/:key", squares.Square) //cached
	r.HandleFunc("/isogrids/:key", isogrids.Isogrids)
	r.HandleFunc("/spaceinvaders/:key", spaceinvaders.SpaceInvaders)

	r.HandleFunc("/labs/checkerboard", checkerboard.Checkerboard)
	r.HandleFunc("/labs/squares/random", squares.Random)
	r.HandleFunc("/labs/isogrids/hexa", isogrids.Hexa)
	r.HandleFunc("/labs/isogrids/hexa/:key", isogrids.Hexa)
	r.HandleFunc("/labs/isogrids/skeleton", isogrids.Skeleton)
	r.HandleFunc("/labs/isogrids/diagonals", isogrids.Diagonals)
	r.HandleFunc("/labs/isogrids/halfdiagonals", isogrids.HalfDiagonals)
	r.HandleFunc("/labs/isogrids/gridbw", isogrids.GridBW)
	r.HandleFunc("/labs/isogrids/labs/grid2colors", isogrids.Grid2Colors)
	r.HandleFunc("/labs/isogrids/random", isogrids.Random)
	r.HandleFunc("/labs/isogrids/random-mirror", isogrids.RandomMirror)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
