package main

func collisionDetection(x1, y1, w1, h1, x2, y2, w2, h2, tolerance int) bool {

	if x1 == x2 && y1 == y2 {
		return true
	}

	return false
}
