package isogrids

func mirrorCoordinates(xs []int, lines, fringeSize, offset int) (xsMirror []int) {

	xsMirror = make([]int, len(xs))
	for i, _ := range xs {
		xsMirror[i] = (lines * fringeSize) - xs[i] + offset
	}
	return
}
