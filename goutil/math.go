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
	if len(a) != len(b) {
		panic("mismatched slices")
	}

	d := float64(0)
	for i := 0; i < len(a); i++ {
		d += math.Pow(float64(a[i]-b[i]), 2)
	}
	return math.Sqrt(d)
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

// raise x to the y
func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
