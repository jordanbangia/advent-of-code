package main

import (
	"fmt"
	"strconv"
	"strings"

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

func solve1(input []string) error {
	horizontal := 0
	depth := 0

	for _, l := range input {
		parts := strings.Split(l, " ")
		incr, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			horizontal += incr
		case "down":
			depth += incr
		case "up":
			depth -= incr
		}
	}

	fmt.Printf("Solution 1: %d\n", horizontal*depth)
	return nil
}

func solve2(input []string) error {
	horizontal := 0
	depth := 0
	aim := 0

	for _, l := range input {
		parts := strings.Split(l, " ")

		incr, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			horizontal += incr
			depth += aim * incr
		case "down":
			aim += incr
		case "up":
			aim -= incr
		}
	}

	fmt.Printf("Solution 2: %d\n", horizontal*depth)
	return nil
}
