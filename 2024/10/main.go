package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			p1, p2,
		),
		false,
	)
}

type StackItem struct {
	i, j, currentNum int
	path             []string
}

var dirs [][]int = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func p1(inputText []string) (int, error) {
	g, trailheads := parseText(inputText)

	totalScore := 0
	for i, trailhead := range trailheads {
		s, _ := walkTrails(trailhead, g)
		fmt.Printf("%d: (%d, %d) - %d\n", i+1, trailhead[0], trailhead[1], s)
		totalScore += s
	}
	return totalScore, nil
}

func p2(inputText []string) (int, error) {
	g, trailheads := parseText(inputText)

	totalRating := 0
	for i, trailhead := range trailheads {
		_, r := walkTrails(trailhead, g)
		fmt.Printf("%d: (%d, %d) - %d\n", i+1, trailhead[0], trailhead[1], r)
		totalRating += r
	}
	return totalRating, nil
}

func parseText(inputText []string) (map[string]int, [][]int) {
	g := map[string]int{}

	trailheads := [][]int{}

	for i, line := range inputText {
		for j, r := range line {
			k := goutil.Key(i, j)
			g[k] = goutil.Atoi(string(r))
			if g[k] == 0 {
				trailheads = append(trailheads, []int{i, j})
			}
		}
	}
	return g, trailheads
}

func walkTrails(trailhead []int, g map[string]int) (int, int) {
	s := goutil.NewStack[StackItem]()
	s.Push(StackItem{
		i:          trailhead[0],
		j:          trailhead[1],
		currentNum: g[goutil.Key(trailhead[0], trailhead[1])],
	})

	reachedEndPositions := map[string]struct{}{}
	pathsTaken := map[string]struct{}{}

	for !s.IsEmpty() {
		x, hasValue := s.Pop()
		if !hasValue {
			break
		}

		for _, d := range dirs {
			next := []int{
				x.i + d[0],
				x.j + d[1],
			}
			k := goutil.Key(next[0], next[1])
			if val, exists := g[k]; exists {
				if val-x.currentNum == 1 {
					// this a move we can make
					if val == 9 {
						// we can end this look up
						reachedEndPositions[k] = struct{}{}
						pathsTaken[strings.Join(append(x.path, k), "")] = struct{}{}
					} else {
						newPath := goutil.Duplicate(x.path)
						newPath = append(newPath, k)
						s.Push(StackItem{
							i:          next[0],
							j:          next[1],
							currentNum: val,
							path:       newPath,
						})
					}
				}
			}
		}
	}
	return len(reachedEndPositions), len(pathsTaken)
}
