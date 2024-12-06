package goutil

import "fmt"

func Key(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}

type Coord struct {
	X, Y int
}

func (c *Coord) Key() string {
	return Key(c.X, c.Y)
}
