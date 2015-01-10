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
	r.HandleFunc("/grid/[0-8]?", gridColorHandler)

	r.HandleFunc("/grid/random/?", gridRandomHandler)
	r.HandleFunc("/grid/random/[0-8]/?", gridRandomColorHandler)

	r.HandleFunc("/grid/random/symetric/x/?", gridRandomSymetricXHandler)
	r.HandleFunc("/grid/random/symetric/x/[0-8]?", gridRandomSymetricXColorHandler)
	r.HandleFunc("/grid/random/symetric/y/?", gridRandomSymetricYHandler)
	r.HandleFunc("/grid/random/symetric/y/[0-8]?", gridRandomSymetricYColorHandler)

	r.HandleFunc("/grid/random/symetric/?", gridRandomSymetricXHandler)

	r.HandleFunc("/gradient/?", gradientHandler)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
