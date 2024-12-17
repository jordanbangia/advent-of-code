package goutil

func CopyMap[K comparable, T any](src map[K]T) map[K]T {
	m := map[K]T{}

	for k, v := range src {
		m[k] = v
	}
	return m
}
