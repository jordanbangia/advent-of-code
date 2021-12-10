package main

import (
	"fmt"
	"strconv"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = solve1(input)
	// _ = solve2(input)
}

func solve1(input []string) error {

	hMap := make([][]int, len(input))

	for i, l := range input {
		row := make([]int, len(l))
		for j, c := range l {
			row[j], _ = strconv.Atoi(string(c))
		}
		hMap[i] = row
	}

	m := len(hMap)
	n := len(hMap[0])

	isLowPoint := func(x, y int) bool {
		comparePoint := hMap[x][y]

		isLowestPoint := true
		if x > 0 {
			isLowestPoint = isLowestPoint && comparePoint < hMap[x-1][y]
		}
		if x < m-1 {
			isLowestPoint = isLowestPoint && comparePoint < hMap[x+1][y]
		}
		if y > 0 {
			isLowestPoint = isLowestPoint && comparePoint < hMap[x][y-1]
		}
		if y < n-1 {
			isLowestPoint = isLowestPoint && comparePoint < hMap[x][y+1]
		}

		return isLowestPoint
	}

	totalRiskLevel := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isLowPoint(i, j) {
				totalRiskLevel += (1 + hMap[i][j])
			}
		}
	}

	fmt.Printf("Solution 1: %d\n", totalRiskLevel)

	return nil
}

func solve2(input []string) error {
	// to find a basin, we need to find
}
