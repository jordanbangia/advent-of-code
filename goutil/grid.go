package goutil

import "fmt"

func Key(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}

// assumes a grid where (0,0) is top left
// positive y goes down
// positive x goes to the right
type Coord struct {
	X, Y int
}

func (c *Coord) Key() string {
	return Key(c.X, c.Y)
}

// Moves this coordinate 1 unit in the d direction
func (c *Coord) Move(d Direction) {
	switch d {
	case Up:
		c.Y -= 1
	case Right:
		c.X += 1
	case Down:
		c.Y += 1
	case Left:
		c.X -= 1
	}
}

// Determines what the next position would look like
// based on moving 1 unit in the d Direction
func (c *Coord) PeakMove(d Direction) *Coord {
	switch d {
	case Up:
		return &Coord{c.X, c.Y - 1}
	case Right:
		return &Coord{c.X + 1, c.Y}
	case Down:
		return &Coord{c.X, c.Y + 1}
	case Left:
		return &Coord{c.X - 1, c.Y}
	}
	return nil
}

func (c *Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type MatrixCoord struct {
	R int
	C int
}

func (m *MatrixCoord) Key() string {
	return Key(m.R, m.C)
}

func (m *MatrixCoord) PeakMove(d Direction) *MatrixCoord {
	switch d {
	case Up:
		return &MatrixCoord{m.R - 1, m.C}
	case Right:
		return &MatrixCoord{m.R, m.C + 1}
	case Down:
		return &MatrixCoord{m.R + 1, m.C}
	case Left:
		return &MatrixCoord{m.R, m.C - 1}
	}
	return nil
}

func (m *MatrixCoord) Move(d Direction) {
	nextMove := m.PeakMove(d)
	m.R, m.C = nextMove.R, nextMove.C
}

func (m *MatrixCoord) String() string {
	return m.Key()
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type DirList struct {
	l    []Direction
	curr int
}

func (d *DirList) PeakNext() Direction {
	return d.l[(d.curr+1)%len(d.l)]
}

func (d *DirList) Next() {
	d.curr = (d.curr + 1) % len(d.l)
}

func (d *DirList) Direction() Direction {
	return d.l[d.curr]
}

func NewDirList(directionOrder []Direction, startDirection Direction) *DirList {
	startIndex := 0
	for i, d := range directionOrder {
		if d == startDirection {
			startIndex = i
			break
		}
	}

	return &DirList{l: directionOrder, curr: startIndex}
}
