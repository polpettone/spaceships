package engine

import "math"

func CollisionDetection(x1, y1, x2, y2, w1, w2, tolerance int) bool {

	d := calcDistance(x1, y1, x2, y2)

	maxW := int(math.Max(float64(w1/2), float64(w2/2)))

	if d <= maxW {
		return true
	}

	return false

}

func calcDistance(x1, y1, x2, y2 int) int {
	d := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
	return int(d)

}
