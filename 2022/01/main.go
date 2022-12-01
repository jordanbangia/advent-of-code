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

	foodStorage := []int{}

	currentFood := 0
	for _, n := range input {
		if n == "" {
			// new line, means we can start a new elf
			foodStorage = append(foodStorage, currentFood)
			currentFood = 0
		} else {
			num, _ := strconv.Atoi(n)
			currentFood += num
		}
	}

	sort.Ints(foodStorage)
	for i, j := 0, len(foodStorage)-1; i < j; i, j = i+1, j-1 {
		foodStorage[i], foodStorage[j] = foodStorage[j], foodStorage[i]
	}

	sol1 := foodStorage[0]
	sol2 := foodStorage[0] + foodStorage[1] + foodStorage[2]
	goutil.PrintSolution(sol1, sol2)
}
