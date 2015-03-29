package format

import "testing"

func TestFormat(t *testing.T) {
	if JPEG != 0 {
		t.Errorf("Expected 0 got,", JPEG)
	}
	if SVG != 1 {
		t.Errorf("Expected 1 got,", SVG)
	}
}
