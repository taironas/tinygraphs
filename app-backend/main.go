package main

import (
	"bytes"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strconv"
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

func blackHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func greenHandler(w http.ResponseWriter, r *http.Request) {
	// prepare
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	green := color.RGBA{0, 128, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("unable to encode image")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Fatalln("unable to write image")
	}
}
