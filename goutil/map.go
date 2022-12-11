package goutil

import "fmt"

func Key(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}
