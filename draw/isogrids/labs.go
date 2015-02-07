package isogrids

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

func Diagonals(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		canvas.Line(x, 0, x, lastY, style)
	}

	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x, 0, style)
		}
	}

	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x, 0, style)
		}
	}

	canvas.End()
}

// HalfDiagonals builds an image with 10x10 grids of half diagonals
func HalfDiagonals(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		canvas.Line(x, 0, x, lastY, style)
	}

	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}

	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

	canvas.End()
}

// Skeleton builds an image with 10x10 grids of half diagonals
func Skeleton(w http.ResponseWriter, key string, color1, color2 color.RGBA, size int) {
	canvas := svg.New(w)
	canvas.Start(size, size)

	lines := 10
	fringeSize := size / lines

	for xL := 0; xL <= lines; xL++ {
		x := xL * fringeSize
		firstY := 0
		lastY := (lines) * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		if (xL % 2) != 0 {
			lastY = lastY - fringeSize/2
			firstY = fringeSize / 2
		}
		canvas.Line(x, firstY, x, lastY, style)
	}

	for xL := 0; xL <= 2*lines; xL++ {
		x := xL * fringeSize
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		xPrev := 0
		yPrev := (xL) * fringeSize
		if yPrev > 0 {
			canvas.Line(xPrev, yPrev, x*2, 0, style)
		}
	}

	for xL := -2 * lines; xL <= 2*lines; xL++ {
		x := xL * fringeSize * 2
		style := fmt.Sprintf("stroke:black;stroke-width:2; %s", draw.FillFromRGBA(color2))
		xPrev := lines * fringeSize
		yPrev := (lines - xL*2) * fringeSize
		if yPrev > 0 {
			canvas.Line(x, 0, xPrev, yPrev/2, style)
		}
	}

	canvas.End()
}
