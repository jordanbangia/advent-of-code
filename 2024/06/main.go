package main

import (
	"fmt"
	"sync"

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

func copyGrid(g map[string]rune) map[string]rune {
	newM := map[string]rune{}
	for k, v := range g {
		newM[k] = v
	}
	return newM
}

func checkForLoop(sR, sC int, sD goutil.Direction, grid map[string]rune) bool {
	visited := map[string]struct{}{}

	curr := &goutil.MatrixCoord{R: sR, C: sC}
	dir := goutil.NewDirList([]goutil.Direction{goutil.Up, goutil.Right, goutil.Down, goutil.Left}, sD)

	k := func() string {
		return fmt.Sprintf("%s_%d", curr.Key(), dir.Direction())
	}

	for {
		if _, exists := visited[k()]; exists {
			// loop detected
			return true
		}
		visited[k()] = struct{}{}

		next := curr.PeakMove(dir.Direction())
		spot, exists := grid[next.Key()]
		if !exists {
			return false
		}

		switch spot {
		case '.':
			curr.Move(dir.Direction())
		case '#':
			dir.Next()
		}
	}
}

func p2(inputText []string) (int, error) {
	// parse the input
	// this time we keep an object at each cell that tracks
	// which directions we have visted this cell in
	// i.e. we went through this cell going up, down, etc.

	grid := map[string]rune{}
	sR, sC, dirStart := -1, -1, goutil.Up
	for r, l := range inputText {
		for col, c := range l {
			isStartCell := c == 'v' || c == '^' || c == '>' || c == '<'
			if isStartCell {
				sR = r
				sC = col
				dirStart = cToDir(c)
				c = '.'
			}
			grid[goutil.Key(r, col)] = c
		}
	}

	// these are all the points that we'll visit along the way
	visitedPath := visitPath(sR, sC, dirStart, grid)

	results := make(chan bool, len(visitedPath))

	waitGroup := &sync.WaitGroup{}

	for visitedPoint := range visitedPath {
		copyGrid := copyGrid(grid)
		if visitedPoint == goutil.Key(sR, sC) {
			continue
		}
		copyGrid[visitedPoint] = '#'
		waitGroup.Add(1)
		go func(vP string, cG map[string]rune) {
			defer waitGroup.Done()
			itsWorks := checkForLoop(sR, sC, dirStart, copyGrid)
			results <- itsWorks
		}(visitedPoint, copyGrid)
	}
	waitGroup.Wait()
	close(results)

	spotsThatWork := 0
	for r := range results {
		if r {
			spotsThatWork += 1
		}
	}

	return spotsThatWork, nil
}
