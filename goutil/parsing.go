package goutil

import "strconv"

func Atoi[T string | rune | byte](s T) int {
	r, _ := strconv.Atoi(string(s))
	return r
}
