package main

import (
	"fmt"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func cToDir(c rune) goutil.Direction {
	switch c {
	case 'v':
		return goutil.Down
	case '>':
		return goutil.Right
	case '<':
		return goutil.Left
	case '^':
		return goutil.Up
	default:
		panic(fmt.Sprintf("unknown direction %s", string(c)))
	}
}

func p1(inputText []string) (int, error) {
	grid := map[string]rune{}
	sR, sC, dir := -1, -1, goutil.Up
	for r, l := range inputText {
		for col, c := range l {
			if c == 'v' || c == '^' || c == '>' || c == '<' {
				sR = r
				sC = col
				dir = cToDir(c)
				println(sR, sC, string(c), dir)
				c = 'x'

			}
			grid[goutil.Key(r, col)] = c
		}
	}
	return len(visitPath(sR, sC, dir, grid)), nil
}

func visitPath(sR, sC int, sD goutil.Direction, grid map[string]rune) map[string]struct{} {
	visited := map[string]struct{}{}
	cCoord := &goutil.MatrixCoord{R: sR, C: sC}
	cDL := goutil.NewDirList([]goutil.Direction{goutil.Up, goutil.Right, goutil.Down, goutil.Left}, sD)
	for {

		// first, attempt to take a step forward
		// check whats going to happen
		nextCord := cCoord.PeakMove(cDL.Direction())
		spot, exists := grid[nextCord.Key()]
		if !exists {
			// the next move would move off the map
			return visited
		}

		switch spot {
		case '.':
			// we can move to the next location, and its unvisited
			// first mark that its been visited
			visited[nextCord.Key()] = struct{}{}
			// then move
			cCoord.Move(cDL.Direction())
		case '#':
			// the next location can't be moved to
			// we need to change direction, and then try the overall loop again
			cDL.Next()
		}
	}
}
