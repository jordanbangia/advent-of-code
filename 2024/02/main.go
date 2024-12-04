package main

import (
	"strconv"
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

func p1(inputText []string) (int, error) {
	safe := 0

	checkLevels := func(levels []int) bool {
		prevDiff := 0
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]

			if diff == 0 {
				// no change, unsafe
				return false
			}
			if goutil.Abs(diff) > 3 {
				// too much change, unsafe
				return false
			}
			if prevDiff != 0 && diff*prevDiff < 0 {
				// sign has changed, unsafe
				return false
			}
			prevDiff = diff
		}
		return true
	}

	for _, line := range inputText {
		levelsStrParts := strings.Split(line, " ")
		levels := make([]int, len(levelsStrParts))
		for i, l := range levelsStrParts {
			levels[i], _ = strconv.Atoi(l)
		}

		if isSafe := checkLevels(levels); isSafe {
			safe += 1
			// fmt.Printf("%+v\n", levels)
		}
	}

	return safe, nil
}

func p2(inputText []string) (int, error) {
	safe := 0

	checkLevels := func(levels []int) bool {
		prevDiff := 0
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]

			if diff == 0 {
				// no change, unsafe
				return false
			}
			if goutil.Abs(diff) > 3 {
				// too much change, unsafe
				return false
			}
			if prevDiff != 0 && diff*prevDiff < 0 {
				// sign has changed, unsafe
				return false
			}
			prevDiff = diff
		}
		return true
	}

	checkLevelsWithOneMissing := func(levels []int) bool {
		for i := 0; i < len(levels); i++ {
			tmp := make([]int, len(levels))
			copy(tmp, levels)
			if isSafe := checkLevels(
				removeIndex(tmp, i),
			); isSafe {
				return true
			}
		}
		return false
	}

	for _, line := range inputText {
		levelsStrParts := strings.Split(line, " ")
		levels := make([]int, len(levelsStrParts))
		for i, l := range levelsStrParts {
			levels[i], _ = strconv.Atoi(l)
		}

		if checkLevels(levels) {
			// safe with no changes
			safe += 1
		} else if checkLevelsWithOneMissing(levels) {
			safe += 1
		}
	}

	return safe, nil
}

func removeIndex(s []int, i int) []int {
	ret := []int{}
	ret = append(ret, s[:i]...)
	ret = append(ret, s[i+1:]...)
	return append(s[:i], s[i+1:]...)
}
