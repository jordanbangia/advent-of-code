package goutil

func ReverseStrArray(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}

func Sum(a []int) int {
	s := 0
	for _, k := range a {
		s += k
	}
	return s
}

func Duplicate[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func DuplicateArrayOfArray[T any](src [][]T) [][]T {
	dst := make([][]T, len(src))
	for i := range src {
		dst[i] = Duplicate(src[i])
	}
	return dst
}
