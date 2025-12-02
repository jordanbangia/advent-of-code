package main

import (
	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.SolutionMain(goutil.NewOneFuncSolution(solStupid))
}

const dialStart = 50

func solStupid(text []string) (int, int, error) {
	place := dialStart

	endZeroes := 0
	crossZeroes := 0

	for _, line := range text {
		dir := line[0]
		num := goutil.Atoi(line[1:])

		step := 1
		if dir == 'L' {
			step = -1
		}

		for range num {
			place += step
			if place < 0 {
				place += 100
			}
			place = place % 100
			if place == 0 {
				crossZeroes += 1
			}
		}
		if place == 0 {
			endZeroes += 1
		}
	}
	return endZeroes, crossZeroes, nil
}
