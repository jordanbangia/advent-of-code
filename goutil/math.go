package goutil

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func Dist(a, b []int) float64 {
	x := (a[0] - b[0]) * (a[0] - b[0])
	y := (a[1] - b[1]) * (a[1] - b[1])
	return math.Sqrt(float64(x + y))
}

func PosMod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}
