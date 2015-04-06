package isogrids

const (
	left = iota
	right
)

// trianglePosition hold the triplet (x, y, direction)
// direction been either 'left' or 'right'
type trianglePosition struct {
	x, y, direction int
}

var (
	// triangles in an array of arrray triangle positions.
	// each array correspond to a triangle, there are 6 of them,
	// indexes from 0 to 5, they form an hexagon.
	// each triangle to composed of 9 subtriangles ordered
	// from up left to down right
	triangles = [][]trianglePosition{
		[]trianglePosition{
			{0, 1, right},
			{0, 2, right},
			{0, 3, right},
			{0, 2, left},
			{0, 3, left},
			{1, 2, right},
			{1, 3, right},
			{1, 2, left},
			{2, 2, right},
		},
		[]trianglePosition{
			{0, 1, left},
			{1, 1, right},
			{1, 0, left},
			{1, 1, left},
			{2, 0, right},
			{2, 1, right},
			{2, 0, left},
			{2, 1, left},
			{2, 2, left},
		}, []trianglePosition{
			{3, 0, right},
			{3, 1, right},
			{3, 2, right},
			{3, 0, left},
			{3, 1, left},
			{4, 0, right},
			{4, 1, right},
			{4, 1, left},
			{5, 1, right},
		},
		[]trianglePosition{
			{3, 2, left},
			{4, 2, right},
			{4, 2, left},
			{4, 3, left},
			{5, 2, right},
			{5, 3, right},
			{5, 1, left},
			{5, 2, left},
			{5, 3, left},
		},
		[]trianglePosition{
			{3, 3, right},
			{3, 4, right},
			{3, 5, right},
			{3, 3, left},
			{3, 4, left},
			{4, 3, right},
			{4, 4, right},
			{4, 4, left},
			{5, 4, right},
		},
		[]trianglePosition{
			{0, 4, left},
			{1, 4, right},
			{1, 3, left},
			{1, 4, left},
			{2, 3, right},
			{2, 4, right},
			{2, 3, left},
			{2, 4, left},
			{2, 5, left},
		},
	}
)

func newTrianglePosition(x, y, d int) trianglePosition {
	return trianglePosition{x, y, d}
}

// isInTriangle tells you whether the triples (x, y,direction)
// is a position inside one of the triangles.
func (tp *trianglePosition) isInTriangle() bool {
	return tp.triangleID() != -1
}

// triangleID returns the triangle id (from 0 to 5)
// that has a match with the position given as param.
// returns -1 if a match is not found.
func (tp *trianglePosition) triangleID() int {

	for i, t := range triangles {
		for _, ti := range t {
			if ti.x == tp.x && ti.y == tp.y && tp.direction == ti.direction {
				return i
			}
		}
	}
	return -1
}

// subTriangleID returns the sub triangle id (from 0 to 8)
// that has a match with the position given as param.
// returns -1 if a match is not found.
func (tp *trianglePosition) subTriangleID() int {

	for _, t := range triangles {
		for i, ti := range t {
			if ti.x == tp.x && ti.y == tp.y && tp.direction == ti.direction {
				return i
			}
		}
	}
	return -1
}

func subTriangleRotations(lookforSubTriangleID int) []int {

	m := map[int][]int{
		0: []int{0, 6, 8, 8, 2, 0},
		1: []int{1, 2, 5, 7, 6, 3},
		2: []int{2, 0, 0, 6, 8, 8},
		3: []int{3, 4, 7, 5, 4, 1},
		4: []int{4, 1, 3, 4, 7, 5},
		5: []int{5, 7, 6, 3, 1, 2},
		6: []int{6, 3, 1, 2, 5, 7},
		7: []int{7, 5, 4, 1, 3, 4},
		8: []int{8, 8, 2, 0, 0, 6},
	}
	if v, ok := m[lookforSubTriangleID]; ok {
		return v
	}
	return nil
}

// rotationId returns the original sub triangle id
// if the current triangle was rotated to position 0.
func (tp *trianglePosition) rotationID() int {
	currentTID := tp.triangleID()
	currentSTID := tp.subTriangleID()
	numberOfSubTriangles := 9
	for i := 0; i < numberOfSubTriangles; i++ {
		rotations := subTriangleRotations(i)
		if rotations[currentTID] == currentSTID {
			return i
		}
	}
	return -1
}
