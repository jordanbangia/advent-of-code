package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, _ := goutil.ReadFile("test.txt")

	col1 := make([]int, len(input))
	col2 := make([]int, len(input))
	col2Count := map[int]int{}
	for i, line := range input {
		parts := strings.Split(line, " ")
		col1[i], _ = strconv.Atoi(parts[0])
		col2Num, _ := strconv.Atoi(parts[3])
		col2[i] = col2Num
		col2Count[col2Num] += 1
	}

	sort.Ints(col1)
	sort.Ints(col2)

	dist := 0
	for i := 0; i < len(col1); i++ {
		dist += goutil.Abs(col1[i] - col2[i])
	}

	similarity := 0
	for _, col1Num := range col1 {
		similarity += col1Num * col2Count[col1Num]
	}

	goutil.PrintSolution(dist, similarity)
}
