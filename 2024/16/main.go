package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewOneFuncSolution(p),
		false,
	)
}

type S struct {
	i, j  int
	score int
	cdi   int

	prev *S

	// we should use a set or something like that
	// but copying the maps around is annoying and slow
	// instead we use a string that looks like
	// x,y|x,y|x,y
	// and we dedup after the fact
	// its a given that the shortest routes will avoid
	// duplicates in the path
	usedTiles string
}

func p(inputText []string) (int, int, error) {
	paths := findPaths(inputText)

	finishedScore := -1
	for _, p := range paths {
		if finishedScore == -1 {
			finishedScore = p.score
		} else {
			finishedScore = goutil.Min(finishedScore, p.score)
		}
	}

	usedTiles := map[string]struct{}{}
	for _, p := range paths {
		if p.score != finishedScore {
			continue
		}
		for _, coord := range strings.Split(p.usedTiles, "|") {
			usedTiles[coord] = struct{}{}
		}
	}

	return finishedScore, len(usedTiles), nil
}

func findPaths(inputText []string) []*S {
	g := map[string]string{}

	startI, startJ := 0, 0

	for i, l := range inputText {
		for j, c := range l {
			if c == 'S' {
				startI = i
				startJ = j
				c = '.'
			}
			g[goutil.Key(i, j)] = string(c)
		}
	}

	dirs := []goutil.Direction{
		goutil.Right,
		goutil.Down,
		goutil.Left,
		goutil.Up,
	}

	stack := goutil.NewQueue[*S]()
	stack.Push(&S{
		i:         startI,
		j:         startJ,
		score:     0,
		cdi:       0,
		usedTiles: fmt.Sprintf("%d,%d", startI, startJ),
	})

	/**
	Problem: we can get into situations where we start looping around in circles forever
	We need a way to trim certain paths.

	1. We have an immediate backtrack check - if we're going to go to the same position again, we don't do that.
	2. We instead could record for each position and in each direction, whats the lowest score that we've seen there
	if we come up on that with a route that its significantly great (>1000 over), then we can trim the path
	as likely we just took a really inefficent spin around the map
	**/

	seenStates := map[string]int{}

	finishedPaths := []*S{}

	for !stack.IsEmpty() {
		c, _ := stack.Pop()

		k := fmt.Sprintf("%d,%d,%d", c.i, c.j, c.cdi)
		if prevScore, hasSeenState := seenStates[k]; hasSeenState && (c.score-prevScore) > 1000 {
			// if we've previously hit this state but its significantly higher now
			// then we can sorta ignore it as we know that we could have gotten here
			// with a lower score
			continue
		} else if !hasSeenState {
			seenStates[k] = c.score
		} else {
			seenStates[k] = goutil.Min(seenStates[k], c.score)
		}

		// from a given position, we have a couple things to try

		// 1. try to move forward.  Double check that its even possible
		dR, dC := goutil.MoveVals(dirs[c.cdi])
		pos, exists := g[goutil.Key(c.i+dR, c.j+dC)]
		if exists && pos == "E" {
			// don't bother doing the other stuff, cause we could do a bunch
			// of turns but it would only drive up the score
			finalizedS := &S{score: c.score + 1, usedTiles: fmt.Sprintf("%s|%d,%d", c.usedTiles, c.i+dR, c.j+dC)}
			finishedPaths = append(finishedPaths, finalizedS)
		} else if exists && pos != "#" {
			stack.Push(&S{
				i:         c.i + dR,
				j:         c.j + dC,
				score:     c.score + 1,
				cdi:       c.cdi,
				usedTiles: fmt.Sprintf("%s|%d,%d", c.usedTiles, c.i+dR, c.j+dC),
			})
		}

		// 2. try to turn 90 degrees clockwise and 90 degrees counter clockwise
		for _, newD := range []int{
			goutil.PosMod((c.cdi + 1), len(dirs)),
			goutil.PosMod((c.cdi - 1), len(dirs)),
		} {
			n := &S{
				i:         c.i,
				j:         c.j,
				score:     c.score + 1000,
				cdi:       newD,
				prev:      c,
				usedTiles: c.usedTiles,
			}
			if c.prev == nil {
				stack.Push(n)
			} else if c.prev.i == c.i && c.prev.j == c.j && c.prev.cdi == newD {
				continue
			}
		}
	}
	return finishedPaths
}
