package main

func appendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1
	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)
	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
	}
	copy(z[len(x):], y)
	return z
}
