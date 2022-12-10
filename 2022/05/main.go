package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(input)
	part2(input)

}

func part1(input []string) {
	stacks := buildStacks()
	for _, line := range input {
		parts := strings.Split(line, " ")
		moveCount := goutil.Atoi(parts[1])
		from := goutil.Atoi(parts[3]) - 1
		to := goutil.Atoi(parts[5]) - 1

		for c := 0; c < moveCount; c++ {
			ele, _ := stacks[from].Pop()
			stacks[to].Push(ele)
		}
	}
	printResponse(stacks)
}

func part2(input []string) {
	stacks := buildStacks()
	for _, line := range input {
		parts := strings.Split(line, " ")
		moveCount := goutil.Atoi(parts[1])
		from := goutil.Atoi(parts[3]) - 1
		to := goutil.Atoi(parts[5]) - 1

		mid := goutil.Stack([]string{})
		for c := 0; c < moveCount; c++ {
			ele, _ := stacks[from].Pop()
			mid.Push(ele)
		}

		for !mid.IsEmpty() {
			ele, _ := mid.Pop()
			stacks[to].Push(ele)
		}
	}
	printResponse(stacks)
}

func printResponse(stacks []goutil.Stack) {
	response := []string{}
	for i := range stacks {
		response = append(response, stacks[i][len(stacks[i])-1])
	}
	fmt.Println(strings.Join(response, ""))
}

func buildStacks() []goutil.Stack {
	return []goutil.Stack{
		goutil.ReverseStrArray([]string{"Q", "G", "P", "R", "L", "C", "T", "F"}),
		goutil.ReverseStrArray([]string{"J", "S", "F", "R", "W", "H", "Q", "N"}),
		goutil.ReverseStrArray([]string{"Q", "M", "P", "W", "H", "B", "F"}),
		goutil.ReverseStrArray([]string{"F", "D", "T", "S", "V"}),
		goutil.ReverseStrArray([]string{"Z", "F", "V", "W", "D", "L", "Q"}),
		goutil.ReverseStrArray([]string{"S", "L", "C", "Z"}),
		goutil.ReverseStrArray([]string{"F", "D", "V", "M", "B", "Z"}),
		goutil.ReverseStrArray([]string{"B", "J", "T"}),
		goutil.ReverseStrArray([]string{"H", "P", "S", "L", "G", "B", "N", "Q"}),
	}
}
