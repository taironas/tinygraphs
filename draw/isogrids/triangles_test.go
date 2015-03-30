package isogrids

import "testing"

func TestDistanceTo3rdPoint(t *testing.T) {

	got := distanceTo3rdPoint(4)
	expected := 4
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}
