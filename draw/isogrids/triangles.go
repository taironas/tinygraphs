package isogrids

import "math"

func distanceTo3rdPoint(AC int) int {
	// distance from center of vector to third point of equilateral triangles
	// ABC triangle, O is the center of AB vector
	// OC = SQRT(AC^2 - AO^2)
	return int(math.Ceil(math.Sqrt((float64(AC) * float64(AC)) - (float64(AC)/float64(2))*(float64(AC)/float64(2)))))

}

// right1stTriangle computes a right oriented triangle '>'
func right1stTriangle(xL, yL, fringeSize, distance int) (x1, y1, x2, y2, x3, y3 int) {
	x1 = xL * distance
	x2 = xL*distance + distance
	x3 = x1
	y1 = yL * fringeSize
	y2 = y1 + fringeSize/2
	y3 = yL*fringeSize + fringeSize
	return
}

// left1stTriangle computes the coordinates of a left oriented triangle '<'
func left1stTriangle(xL, yL, fringeSize, distance int) (x1, y1, x2, y2, x3, y3 int) {
	x1 = xL*distance + distance
	x2 = xL * distance
	x3 = x1
	y1 = yL * fringeSize
	y2 = y1 + fringeSize/2
	y3 = yL*fringeSize + fringeSize
	return
}

// left2ndTriangle computes the coordinates of a left oriented triangle '<'
func left2ndTriangle(xL, yL, fringeSize, distance int) (x1, y1, x2, y2, x3, y3 int) {
	x1 = xL*distance + distance
	x2 = xL * distance
	x3 = x1
	y1 = yL*fringeSize + fringeSize/2
	y2 = y1 + fringeSize/2
	y3 = yL*fringeSize + fringeSize + fringeSize/2
	return
}

// right2ndTriangle computes the coordinates of a right oriented triangle '>'
func right2ndTriangle(xL, yL, fringeSize, distance int) (x1, y1, x2, y2, x3, y3 int) {
	x1 = xL * distance
	x2 = xL*distance + distance
	x3 = x1
	y1 = yL*fringeSize + fringeSize/2
	y2 = yL + fringeSize
	y3 = yL*fringeSize + fringeSize/2 + fringeSize
	return
}
