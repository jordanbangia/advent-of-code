package main

import (
	"fmt"
	"unicode"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	nums := []int{}

	for _, line := range input {
		first := '0'
		last := '0'
		for _, c := range line {
			if unicode.IsDigit(c) {
				if first == '0' {
					first = c
				}
				last = c
			}
		}

		nums = append(nums, goutil.Atoi(string([]rune{first, last})))
	}

	nums2 := []int{}
	for _, line := range input {
		nums2 = append(nums2, lineToNum(line))
	}

	fmt.Printf("%v", nums2)

	goutil.PrintSolution(goutil.Sum(nums), goutil.Sum(nums2))
}

// 55680

func lineToNum(l string) int {
	// couple rules we can follow
	// we can always replace 'four', 'five', 'six'
	// the stuff we need to be careful about
	// oneight should be one
	// twoone should be two
	// threeight should be three
	// fiveight should be five
	// sevenine should be seven
	// eightwo should be eight
	// eighthree should be eight
	// nineight should be nine

	return goutil.Atoi(string([]rune{'i', 'f'}))
}
