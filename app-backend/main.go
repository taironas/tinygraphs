package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var root = flag.String("root", "app", "khkjh") //"$GOPATH/src/github.com/taironas/greentros/app", "file system path")

func main() {
	//http.HandleFunc("/", hello)
	http.Handle("/", http.FileServer(http.Dir(*root))) //http.FileServer(http.Dir("static")))
	log.Println("location on " + http.Dir(*root))
	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, greentros ngg")
}
