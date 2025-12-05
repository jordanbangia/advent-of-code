package main

import (
	"slices"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.SolutionMain(goutil.NewSolution(part1, part2))
}

func part1(inputText []string) (int, error) {
	mode := "ids"

	freshIngredients := [][]int{}

	isIngredientFresh := func(ingredientID int) bool {
		for _, ingredientRange := range freshIngredients {
			start := ingredientRange[0] - ingredientID
			end := ingredientRange[1] - ingredientID
			if start == 0 || end == 0 || (start < 0 && end > 0) {
				return true
			}
		}
		return false
	}

	freshIngredientsCount := 0

	for _, l := range inputText {
		if mode == "ids" {
			if l == "" || l == " " {
				mode = "available"
				continue
			}
			r := strings.Split(l, "-")
			freshIngredients = append(freshIngredients, []int{goutil.Atoi(r[0]), goutil.Atoi(r[1])})
			// println(freshIngredients[len(freshIngredients)-1][0])
		} else {
			ingredientIsFresh := isIngredientFresh(goutil.Atoi(l))
			// println(goutil.Atoi(l), ingredientIsFresh)
			if ingredientIsFresh {
				freshIngredientsCount++
			}
		}
	}
	return freshIngredientsCount, nil
}

func part2(inputText []string) (int, error) {
	freshIngredients := [][]int{}

	for _, l := range inputText {
		if l == "" || l == " " {
			break
		}
		r := strings.Split(l, "-")
		freshIngredients = append(freshIngredients, []int{goutil.Atoi(r[0]), goutil.Atoi(r[1])})
	}

	slices.SortStableFunc(freshIngredients, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	mergedRanges := [][]int{
		freshIngredients[0],
	}

	for i := 1; i < len(freshIngredients); i++ {
		p := len(mergedRanges) - 1
		nR := freshIngredients[i]

		// nextRange start <= currentRange end
		if nR[0] <= mergedRanges[p][1] || nR[0] == mergedRanges[p][1]+1 {
			// merge
			// take the max as the incoming range could be smaller than the
			// current one - we just know for sure that the incoming range
			// overlaps this range in some capacity
			mergedRanges[p][1] = max(nR[1], mergedRanges[p][1])
		} else {
			// can't merge
			mergedRanges = append(mergedRanges, nR)
		}
	}

	totalFreshIngredients := 0
	for _, mr := range mergedRanges {
		totalFreshIngredients += (mr[1] - mr[0] + 1)
	}

	return totalFreshIngredients, nil
}
