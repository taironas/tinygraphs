package format

// Format type list some image formats.
type Format int

const (
	// JPEG format use it to build jpeg images.
	JPEG Format = iota
	// SVG format use it to build svg images.
	SVG
)
