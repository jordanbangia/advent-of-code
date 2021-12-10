package main

import (
	"fmt"
	"sort"
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
	_ = solve2(input)
}

func parseHeightMap(input []string) [][]int {
	hMap := make([][]int, len(input))

	for i, l := range input {
		row := make([]int, len(l))
		for j, c := range l {
			row[j], _ = strconv.Atoi(string(c))
		}
		hMap[i] = row
	}
	return hMap
}

func solve1(input []string) error {
	hMap := parseHeightMap(input)

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
	hMap := parseHeightMap(input)

	m := len(hMap)
	n := len(hMap[0])

	// 0 is unvisited
	// 1 is visited
	// 9 is impassible

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if hMap[i][j] != 9 {
				hMap[i][j] = 0
			}
		}
	}

	countBasin := func(i, j int) int {
		q := [][]int{[]int{i, j}}

		basinSize := 0
		for len(q) > 0 {
			e := q[0]
			x, y := e[0], e[1]
			q = q[1:]

			if hMap[x][y] != 0 {
				continue
			}

			basinSize += 1
			hMap[x][y] = 1

			if x > 0 {
				q = append(q, []int{x - 1, y})
			}
			if x < m-1 {
				q = append(q, []int{x + 1, y})
			}
			if y > 0 {
				q = append(q, []int{x, y - 1})
			}
			if y < n-1 {
				q = append(q, []int{x, y + 1})
			}
		}
		return basinSize
	}

	basinSizes := []int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if hMap[i][j] == 0 {
				basin := countBasin(i, j)
				basinSizes = append(basinSizes, basin)
			}
		}
	}

	sort.Ints(basinSizes)
	fmt.Println(basinSizes)

	k := len(basinSizes)

	largestBasins := basinSizes[k-1] * basinSizes[k-2] * basinSizes[k-3]

	fmt.Printf("Solution 2: %d\n", largestBasins)
	return nil
}
