package main

import (
	"fmt"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sol1 := 0
	for _, rucksack := range input {
		size := len(rucksack)
		compartment1 := map[rune]struct{}{}

		for i, item := range rucksack {
			if i < size/2 {
				compartment1[item] = struct{}{}
			} else {
				_, hasItem := compartment1[item]
				if hasItem {
					sol1 += calcPrio(item)
					break
				}
			}
		}
	}

	sol2 := 0
	for i := 0; i < len(input); i += 3 {
		pack1 := input[0+i]
		pack2 := input[1+i]
		pack3 := input[2+i]

		p := checkGroup(pack1, pack2, pack3)
		if p == -1 {
			fmt.Println("something has gone wrong with:")
			fmt.Println(pack1)
			fmt.Println(pack2)
			fmt.Println(pack3)
		} else {
			sol2 += p
		}

	}

	goutil.PrintSolution(sol1, sol2)
}

func checkGroup(pack1, pack2, pack3 string) int {
	seenItems := map[rune]int{}

	for _, item := range pack1 {
		seenItems[item] = 1
	}

	for _, item := range pack2 {
		currCount, hasItem := seenItems[item]
		if hasItem && currCount == 1 {
			seenItems[item] = 2
		}
	}

	for _, item := range pack3 {
		currentCount, hasItem := seenItems[item]
		if hasItem && currentCount == 2 {
			return calcPrio(item)
		}
	}
	return -1
}

func calcPrio(item rune) int {
	p := (int(item) - int('a')) + 1
	if p < 0 {
		p = (int(item) - int('A')) + 27
	}
	return p
}
