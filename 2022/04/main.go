package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sol1 := 0
	sol2 := 0
	for _, line := range input {
		sections := strings.Split(line, ",")

		first := sections[0]
		second := sections[1]

		firstStart, firstEnd := parseRange(first)
		secondStart, secondEnd := parseRange(second)

		// check for overlap
		// overlap occurs if
		if rangesCompletelyOverlap([]int{firstStart, firstEnd}, []int{secondStart, secondEnd}) {
			sol1 += 1
		}
		if rangesOverlap([]int{firstStart, firstEnd}, []int{secondStart, secondEnd}) {
			sol2 += 1
		}
	}

	goutil.PrintSolution(sol1, sol2)
}

func parseRange(r string) (int, int) {
	splitted := strings.Split(r, "-")
	return goutil.Atoi(splitted[0]), goutil.Atoi(splitted[1])
}

func rangesCompletelyOverlap(first, second []int) bool {
	return (second[0] >= first[0] && second[1] <= first[1]) || (first[0] >= second[0] && first[1] <= second[1])
}

func rangesOverlap(first, second []int) bool {
	return goutil.Max(first[0], second[0]) <= goutil.Min(first[1], second[1])
}
