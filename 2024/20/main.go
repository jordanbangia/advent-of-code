package main

import (
	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			p1, nil,
		),
		true,
	)
}

const (
	NOT = iota
	AVAILABLE
	CHEATING
	DONE
)

type cheatState struct {
	state      int
	sI, sJ     int
	eI, eJ     int
	cheatSteps int
}

func (c cheatState) nextSteps() int {
	if c.state == CHEATING {
		return c.cheatSteps + 1
	}
	return c.cheatSteps
}

type e struct {
	i, j, steps  int
	prevI, prevJ int // avoid backtracking

	cs cheatState
}

func p1(inputText []string) (int, error) {
	return doIt(inputText, 2, 1)
}

func p2(inputText []string) (int, error) {
	return doIt(inputText, 20, 1)
}

/*
*
what if instead we first tracked out what the path is, and how
far from a given point int he path it is to the end

we could then say:
for each point in the path
- try to cheat by moving through up to 20 spaces of walls
- as soon as you stop cheating cause you hit a path point, calculate
*
*/
func doIt(inputText []string, cheatLength int, mustSaveTime int) (int, error) {
	grid := map[string]int{}

	start := ""
	end := ""
	pathLength := 0
	for i, line := range inputText {
		for j, c := range line {
			if c == 'S' {
				start = goutil.Key(i, j)
				c = '.'
			} else if c == 'E' {
				end = goutil.Key(i, j)
				c = '.'
			}
			if c == '.' {
				grid[goutil.Key(i, j)] = 0
				pathLength += 1
			} else {
				grid[goutil.Key(i, j)] = -1
			}

		}
	}

	// we know that there is a singular path
	// so we can calculate that the max time it would
	// take is just walking this path
	pathLength -= 1

	// the next thing we want to do though is seed the
	// the grid with the "index" of the steps around the path
	// that way we know if path length = x, and some index is x-n
	// that there are n more steps along the path

	fillGridWithPathIndicesAndFindPath := func() [][]int {
		pI, pJ := -1, -1
		cI, cJ := goutil.SplitKey(start)
		endI, endJ := goutil.SplitKey(end)
		stepsRemaining := pathLength
		grid[start] = stepsRemaining
		stepsRemaining -= 1

		path := [][]int{}
		for {
			path = append(path, []int{cI, cJ})
			if cI == endI && cJ == endJ {
				grid[end] = 0
				return path
			}
			for _, dir := range [][]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			} {
				nI, nJ := cI+dir[0], cJ+dir[1]
				if nI == pI && nJ == pJ {
					// avoid backtracking
					continue
				}
				nKey := goutil.Key(nI, nJ)
				if cell, exists := grid[nKey]; !exists || cell == -1 {
					continue
				}
				grid[nKey] = stepsRemaining
				stepsRemaining -= 1
				pI, pJ = cI, cJ
				cI, cJ = nI, nJ
				break
			}
		}
	}
	path := fillGridWithPathIndicesAndFindPath()

	findCheats := func(startPoint []int) {
		// we want to find everysquare thats within cheatLength
		// space of the start point

		for i := 0; i < cheatLength; i++ {
			for j := 0; j < i; j++ {

			}
		}
	}

	cheats := 0
	for _, p := range path {
		cheats += findCheats(p)
	}

	return -1, nil
}
