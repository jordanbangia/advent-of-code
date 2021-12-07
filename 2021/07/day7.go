package main

import (
	"fmt"
	"math"
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
	positions := []int{}

	min := -1
	max := 0
	for _, p := range strings.Split(input[0], ",") {
		i, _ := strconv.Atoi(p)
		positions = append(positions, i)

		if min == -1 {
			min = i
		} else if i < min {
			min = i
		}

		if i > max {
			max = i
		}
	}

	minConsumption := -1
	for i := min; i < max; i++ {
		f := fuelConsumptionSimple(i, positions)
		if minConsumption == -1 {
			minConsumption = f
		} else if f < minConsumption {
			minConsumption = f
		}
	}

	fmt.Printf("Solution 1: %d\n", minConsumption)
	return nil
}

func fuelConsumptionSimple(i int, p []int) int {
	c := 0
	for _, pos := range p {
		c += int(math.Abs(float64(pos - i)))
	}
	return c
}

func fuelConsumptionComplex(i int, p []int) int {
	c := 0
	for _, pos := range p {
		diff := int(math.Abs(float64(pos - i)))
		c += consumptionFromDiff(diff)
	}
	return c
}

var consumptionMap = map[int]int{}

func consumptionFromDiff(diff int) int {
	if _, ok := consumptionMap[diff]; !ok {
		consumptionMap[diff] = diff + consumptionFromDiff(diff-1)
	}
	return consumptionMap[diff]
}

func solve2(input []string) error {
	consumptionMap[0] = 0
	consumptionMap[1] = 1
	consumptionMap[2] = 3
	consumptionMap[3] = 6

	positions := []int{}

	min := -1
	max := 0
	for _, p := range strings.Split(input[0], ",") {
		i, _ := strconv.Atoi(p)
		positions = append(positions, i)

		if min == -1 {
			min = i
		} else if i < min {
			min = i
		}

		if i > max {
			max = i
		}
	}

	minConsumption := -1
	for i := min; i < max; i++ {
		f := fuelConsumptionComplex(i, positions)
		if minConsumption == -1 {
			minConsumption = f
		} else if f < minConsumption {
			minConsumption = f
		}
	}

	fmt.Printf("Solution 2: %d\n", minConsumption)
	return nil
}
