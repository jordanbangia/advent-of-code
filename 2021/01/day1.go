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

	inputNums := make([]int, len(input))
	for i, n := range input {
		num, _ := strconv.Atoi(n)
		inputNums[i] = num
	}

	_ = solve1(inputNums)

	_ = solve2(inputNums)
}

func solve1(input []int) error {

	increases := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	fmt.Printf("Solution 1: %d\n", increases)
	return nil
}

func solve2(input []int) error {
	increases := 0

	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			increases += 1
		}
	}

	fmt.Printf("Solution 2: %d\n", increases)
	return nil
}
