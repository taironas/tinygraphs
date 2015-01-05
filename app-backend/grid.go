package main

import (
	"image"
	"image/color"
	"strconv"
	"strings"
	//	"image/draw"
	"log"
	"net/http"
)

// parse permalink id from URL  and return it
func PermalinkID(r *http.Request, level int64) (int64, error) {
	url := strings.Replace(r.URL.String(), "http://", "", 1)
	path := strings.Split(url, "/")
	// if url has params extract id until the ? character
	var strID string

	strID = path[level]
	intID, err := strconv.ParseInt(strID, 0, 64)
	if err != nil {
		// only try to extract id if were are unable to exracted using the level.
		if strings.Contains(r.URL.String(), "?") {
			strPath := path[level]
			strID = strPath[0:strings.Index(strPath, "?")]
		} else {
			strID = path[level]
		}
		intID, err = strconv.ParseInt(strID, 0, 64)
		if err != nil {
			log.Printf("error when calling PermalinkID with %v.Error: %v", path[level], err)
		}
	}
	return intID, err
}

func colorGridHandler(w http.ResponseWriter, r *http.Request) {
	intID, err := PermalinkID(r, 1)
	if err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		m := image.NewRGBA(image.Rect(0, 0, 240, 240))
		var c color.RGBA
		if intID == 1 {
			c = color.RGBA{
				uint8(0),
				uint8(144),
				167,
				255}
		}
		if intID == 2 {
			c = color.RGBA{
				uint8(232),
				uint8(70),
				134,
				255}
		}
		if intID == 3 {
			c = color.RGBA{
				uint8(255),
				uint8(237),
				81,
				255}
		}
		if intID == 4 {
			c = color.RGBA{
				uint8(177),
				uint8(192),
				24,
				255}
		}
		if intID == 5 {
			c = color.RGBA{
				uint8(208),
				uint8(2),
				120,
				255}
		}
		if intID == 6 {
			c = color.RGBA{
				uint8(0),
				uint8(44),
				47,
				255}
		}
		if intID == 7 {
			c = color.RGBA{
				uint8(29),
				uint8(24),
				18,
				255}
		}
		if intID == 8 {
			c = color.RGBA{
				uint8(33),
				uint8(30),
				26,
				255}
		}

		drawGrid6X6(m, c)
		var img image.Image = m
		writeImage(w, &img)
	}
}

func grid6X6Handler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	color := color.RGBA{
		uint8(255),
		uint8(255),
		255,
		255}
	drawGrid6X6(m, color)
	var img image.Image = m
	writeImage(w, &img)
}

func gradientHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	drawGradient(m)
	var img image.Image = m
	writeImage(w, &img)
}

func drawGrid6X6(m *image.RGBA, color color.RGBA) {
	size := m.Bounds().Size()
	quad := size.X / 6
	for x := 0; x < size.X; x++ {
		val := ((x / quad) % quad) % 2
		for y := 0; y < size.Y; y++ {
			val2 := ((y / quad) % quad) % 2
			if val+val2 == 1 {
				m.Set(x, y, color)
			}
		}
	}
}

func drawGradient(m *image.RGBA) {
	size := m.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255}
			m.Set(x, y, color)
		}
	}
}
