package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			nil, p2,
		),
		false,
	)
}

func addToDict(d map[int][]int, k, v int) {
	if _, exists := d[k]; !exists {
		d[k] = []int{}
	}
	d[k] = append(d[k], v)
}

func p1(inputText []string) (int, error) {
	g := map[string]string{}

	isChar := func(r, c int, t string) bool {
		char, exists := g[goutil.Key(r, c)]
		return exists && char == t
	}

	isWall := func(r, c int) bool { return isChar(r, c, "#") }
	isBox := func(r, c int) bool { return isChar(r, c, "O") }

	startPosition := []int{}

	moves := []goutil.Direction{}

	// maxR := len(inputText)
	// maxC := len(inputText[0])
	for i, l := range inputText {
		for j, c := range l {
			switch c {
			case '^':
				moves = append(moves, goutil.Up)
			case '>':
				moves = append(moves, goutil.Right)
			case '<':
				moves = append(moves, goutil.Left)
			case 'v':
				moves = append(moves, goutil.Down)
			case '#':
				fallthrough
			case '.':
				fallthrough
			case 'O':
				g[goutil.Key(i, j)] = string(c)
			case '@':
				startPosition = []int{i, j}
				g[goutil.Key(i, j)] = "."
			}
		}
	}

	// printGrid(g, maxR, maxC)

	determineBoxesPushed := func(sR, sC, dR, dC int) map[string]string {
		// returns a map of box at initial position -> box at new position
		cR, cC := sR, sC

		updates := map[string]string{}

		for {
			k := goutil.Key(cR, cC)
			if isWall(cR, cC) {
				// we hit a wall,  which means that we wouldn't be able to make this move
				return map[string]string{}
			}

			if isBox(cR, cC) {
				// there's a box, so we should try moving it in the direction
				updates[k] = goutil.Key(cR+dR, cC+dC)
				cR += dR
				cC += dC
				continue
			}

			// its neither a wall nor a box
			// that means we've hit an empty spot
			// so we can return the updates
			return updates
		}
	}

	currentPosition := startPosition
	for _, move := range moves {
		dR, dC := goutil.MoveVals(move)
		np := []int{
			currentPosition[0] + dR,
			currentPosition[1] + dC,
		}

		// if we would move into a wall, we can't do the move
		if isWall(np[0], np[1]) {
			// so skip this move and continue
			continue
		}

		if !isBox(np[0], np[1]) {
			// there isn't a box at the next position
			// so we wouldn't move anything
			// we can just update our current position
			currentPosition = np
			continue
		}

		// the fun part
		// the next position has a box, so we should attempt
		// to move the stack of boxes
		boxPositionUpdates := determineBoxesPushed(np[0], np[1], dR, dC)
		if len(boxPositionUpdates) == 0 {
			// we were not able to move the boxes
			// likely cause a wall was eventually in the way
			// so we shouldn't update our position
			continue
		} else {
			// we were able to make some updates
			// so apply the updates
			for currentBox := range boxPositionUpdates {
				g[currentBox] = "."
			}
			for _, nextBox := range boxPositionUpdates {
				g[nextBox] = "O"
			}
			// then we move our position
			currentPosition = np
		}
	}

	// printGrid(g, maxR, maxC)

	totalCoords := 0
	for position, char := range g {
		if char == "O" {
			r, c := goutil.SplitKey(position)
			totalCoords += (r*100 + c)
		}
	}

	return totalCoords, nil
}

func p2(inputText []string) (int, error) {
	g := map[string]string{}

	isChar := func(r, c int, t string) bool {
		char, exists := g[goutil.Key(r, c)]
		return exists && char == t
	}

	isWall := func(r, c int) bool { return isChar(r, c, "#") }
	isBox := func(r, c int) bool { return isChar(r, c, "[") || isChar(r, c, "]") }

	startPosition := []int{}

	moves := []goutil.Direction{}

	maxR := len(inputText)
	maxC := len(inputText[0]) * 2
	for i, l := range inputText {
		for j, c := range l {
			switch c {
			case '^':
				moves = append(moves, goutil.Up)
			case '>':
				moves = append(moves, goutil.Right)
			case '<':
				moves = append(moves, goutil.Left)
			case 'v':
				moves = append(moves, goutil.Down)
			case '#':
				g[goutil.Key(i, 2*j)] = "#"
				g[goutil.Key(i, 2*j+1)] = "#"
			case '.':
				g[goutil.Key(i, 2*j)] = "."
				g[goutil.Key(i, 2*j+1)] = "."
			case 'O':
				g[goutil.Key(i, 2*j)] = "["
				g[goutil.Key(i, 2*j+1)] = "]"
			case '@':
				startPosition = []int{i, 2 * j}
				g[goutil.Key(i, 2*j)] = "."
				g[goutil.Key(i, 2*j+1)] = "."
			}
		}
	}

	fmt.Println(startPosition)
	printGrid(g, maxR, maxC)

	determineBoxesPushed := func(sR, sC, dR, dC int) (map[string]struct{}, map[string]string) {
		// returns a map of box at initial position -> box at new position

		// this function now tries to move from (sR, sC) in the (dR, dC) direction
		// whenever it hits a box, it expands the things it checks to include the entire
		// box.  It keeps doing this until every possible point moves to an empty space
		// or any possible point hits a wall
		updates := map[string]string{}
		deletes := map[string]struct{}{}

		checks := [][]int{{sR, sC}}
		for {
			boxMoved := false
			nextChecks := [][]int{}
			for _, check := range checks {
				cR, cC := check[0], check[1]

				k := goutil.Key(cR, cC)

				if isWall(cR, cC) {
					// bail out, we hit a wall
					return map[string]struct{}{}, map[string]string{}
				}

				if !isBox(cR, cC) {
					continue
				}

				boxMoved = true
				// we have hit a box
				if g[k] == "[" {
					deletes[goutil.Key(cR, cC)] = struct{}{}
					deletes[goutil.Key(cR, cC+1)] = struct{}{}
					updates[goutil.Key(cR+dR, cC+dC)] = "["
					updates[goutil.Key(cR+dR, cC+1+dC)] = "]"

					if dC == 0 {
						nextChecks = append(nextChecks, []int{cR + dR, cC + dC})
						nextChecks = append(nextChecks, []int{cR + dR, cC + dC + 1})
					} else {
						// we must be travelling right
						// so we can skip to d
						nextChecks = append(nextChecks, []int{cR + dR, cC + dC + 1})
					}
				} else { // g[k] == "]"
					deletes[goutil.Key(cR, cC)] = struct{}{}
					deletes[goutil.Key(cR, cC-1)] = struct{}{}
					updates[goutil.Key(cR+dR, cC+dC)] = "]"
					updates[goutil.Key(cR+dR, cC+dC-1)] = "["

					if dC == 0 {
						nextChecks = append(nextChecks, []int{cR + dR, cC + dC})
						nextChecks = append(nextChecks, []int{cR + dR, cC + dC - 1})
					} else {
						nextChecks = append(nextChecks, []int{cR + dR, cC + dC - 1})
					}
				}
			}

			if !boxMoved {
				return deletes, updates
			}
			checks = nextChecks
		}
	}

	currentPosition := startPosition
	for _, move := range moves {
		dR, dC := goutil.MoveVals(move)
		np := []int{
			currentPosition[0] + dR,
			currentPosition[1] + dC,
		}

		// if we would move into a wall, we can't do the move
		if isWall(np[0], np[1]) {
			// so skip this move and continue
			continue
		}

		if !isBox(np[0], np[1]) {
			// there isn't a box at the next position
			// so we wouldn't move anything
			// we can just update our current position
			currentPosition = np
			continue
		}

		// the fun part
		// the next position has a box, so we should attempt
		// to move the stack of boxes
		deletes, boxPositionUpdates := determineBoxesPushed(np[0], np[1], dR, dC)
		if len(boxPositionUpdates) == 0 {
			// we were not able to move the boxes
			// likely cause a wall was eventually in the way
			// so we shouldn't update our position
			continue
		} else {
			// we were able to make some updates
			// so apply the updates
			for currentBox := range deletes {
				g[currentBox] = "."
			}
			for nextBox, v := range boxPositionUpdates {
				g[nextBox] = v
			}
			// then we move our position
			currentPosition = np
		}
	}

	printGrid(g, maxR, maxC)

	totalCoords := 0
	for position, char := range g {
		if char == "[" {
			r, c := goutil.SplitKey(position)
			totalCoords += (r*100 + c)
		}
	}

	return totalCoords, nil
}

func printGrid(g map[string]string, maxR, maxC int) {
	for r := 0; r < maxR; r++ {
		sb := strings.Builder{}
		for c := 0; c < maxC; c++ {
			sb.WriteString(g[goutil.Key(r, c)])
		}
		fmt.Println(sb.String())
	}
}
