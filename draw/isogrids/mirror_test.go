package isogrids

import "testing"

func TestmirrorCoordinates(t *testing.T) {
	t.Parallel()
	xs := []int{1, 2, 3}
	lines := 10
	fringeSize := 2
	offset := 1
	xsMirror := mirrorCoordinates(xs, lines, fringeSize, offset)
	expected := []int{20, 19, 18}
	if len(expected) != len(xsMirror) {
		t.Errorf("lengths are different")
	}
	for k, _ := range expected {
		if expected[k] != xsMirror[k] {
			t.Errorf("expected %v got %v", expected[k], xsMirror[k])
		}
	}
}
