package isogrids

import "testing"

func TestDistanceTo3rdPoint(t *testing.T) {

	got := distanceTo3rdPoint(4)
	expected := 4
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestRight1stTriangle(t *testing.T) {

	x1, y1, x2, y2, x3, y3 := right1stTriangle(0, 0, 2, 4)
	got := []int{x1, y1, x2, y2, x3, y3}
	expected := []int{0, 0, 4, 1, 0, 2}

	if len(got) != len(expected) {
		t.Errorf("lengths of arrays should be equal")
	}

	for k, _ := range got {
		if got[k] != expected[k] {
			t.Errorf("expected %d got %d", expected[k], got[k])
		}
	}
}

func TestLeft1stTriangle(t *testing.T) {

	x1, y1, x2, y2, x3, y3 := left1stTriangle(0, 0, 2, 4)
	got := []int{x1, y1, x2, y2, x3, y3}
	expected := []int{4, 0, 0, 1, 4, 2}

	if len(got) != len(expected) {
		t.Errorf("lengths of arrays should be equal")
	}

	for k, _ := range got {
		if got[k] != expected[k] {
			t.Errorf("expected %d got %d on element #%d", expected[k], got[k], k)
		}
	}
}

func Testleft2ndTriangle(t *testing.T) {

	x1, y1, x2, y2, x3, y3 := left2ndTriangle(0, 0, 2, 4)
	got := []int{x1, y1, x2, y2, x3, y3}
	expected := []int{4, 1, 0, 2, 4, 3}

	if len(got) != len(expected) {
		t.Errorf("lengths of arrays should be equal")
	}

	for k, _ := range got {
		if got[k] != expected[k] {
			t.Errorf("expected %d got %d on element #%d", expected[k], got[k], k)
		}
	}
}

func TestRight2ndTriangle(t *testing.T) {

	x1, y1, x2, y2, x3, y3 := right2ndTriangle(0, 0, 2, 4)
	got := []int{x1, y1, x2, y2, x3, y3}
	expected := []int{0, 1, 4, 2, 0, 3}

	if len(got) != len(expected) {
		t.Errorf("lengths of arrays should be equal")
	}

	for k, _ := range got {
		if got[k] != expected[k] {
			t.Errorf("expected %d got %d on element #%d", expected[k], got[k], k)
		}
	}
}
