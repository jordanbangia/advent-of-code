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
	fish := []int{}

	for _, l := range input {
		fishStr := strings.Split(l, ",")
		for _, s := range fishStr {
			f, _ := strconv.Atoi(s)
			fish = append(fish, f)
		}
	}

	for i := 0; i < 80; i++ {
		newFish := []int{}
		for f := range fish {
			fish[f] -= 1
			if fish[f] < 0 {
				fish[f] = 6
				newFish = append(newFish, 8)
			}
		}
		fish = append(fish, newFish...)
	}

	fmt.Printf("Solution 1: %d\n", len(fish))
	return nil
}

func solve2(input []string) error {
	// a list works for 80 days as the number doesn't grow that large
	// but gets out of hand for 256 days
	fish := map[int]int{}

	for _, l := range input {
		fishStr := strings.Split(l, ",")
		for _, s := range fishStr {
			f, _ := strconv.Atoi(s)
			fish[f] += 1
		}
	}

	for i := 0; i < 256; i++ {
		updates := map[int]int{}
		for interval, count := range fish {
			if interval == 0 {
				updates[8] += count
				updates[6] += count
			} else {
				updates[interval-1] += count
			}
		}
		fish = updates
	}

	fishAlive := 0
	for _, c := range fish {
		fishAlive += c
	}

	fmt.Printf("Solution 2: %d\n", fishAlive)
	return nil
}
