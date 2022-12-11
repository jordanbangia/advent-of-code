package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(executeMovesForOneKnot(input))
	fmt.Println(executeMovesForTenKnots(input))
}

func executeMovesForOneKnot(input []string) int {
	touched := map[string]bool{}

	head := []int{0, 0}
	tail := []int{0, 0}

	for _, line := range input {
		direction, distance := parseLine(line)
		fmt.Println("executing", direction, distance)
		for i := 0; i < distance; i++ {
			head = updateHead(head, direction)
			tail = updateTail(head, tail)
			touched[goutil.Key(tail[0], tail[1])] = true
		}
	}

	return len(touched)
}

func executeMovesForTenKnots(input []string) int {
	knots := [][]int{
		{0, 0}, // H
		{0, 0}, // 1
		{0, 0}, // 2
		{0, 0}, // 3
		{0, 0}, // 4
		{0, 0}, // 5
		{0, 0}, // 6
		{0, 0}, // 7
		{0, 0}, // 8
		{0, 0}, // 9
	}

	touched := map[string]bool{}

	for _, line := range input {
		direction, distance := parseLine(line)
		fmt.Println("executing", direction, distance)

		for i := 0; i < distance; i++ {
			knots[0] = updateHead(knots[0], direction)
			for x := 1; x < len(knots); x++ {
				knots[x] = updateTail(knots[x-1], knots[x])
			}
			touched[goutil.Key(knots[9][0], knots[9][1])] = true
		}
	}

	return len(touched)
}

func updateHead(head []int, direction string) []int {
	switch direction {
	case "R":
		return []int{head[0] + 1, head[1]}
	case "L":
		return []int{head[0] - 1, head[1]}
	case "U":
		return []int{head[0], head[1] + 1}
	case "D":
		return []int{head[0], head[1] - 1}
	}
	return head
}

func updateTail(head, tail []int) []int {
	if areCloseEnough(head, tail) {
		return []int{tail[0], tail[1]}
	}

	// not close enough, need to update the tail position
	if head[0] == tail[0] {
		if head[1] > tail[1] {
			return []int{tail[0], tail[1] + 1}
		} else {
			return []int{tail[0], tail[1] - 1}
		}
	} else if head[1] == tail[1] {
		if head[0] > tail[0] {
			return []int{tail[0] + 1, tail[1]}
		} else {
			return []int{tail[0] - 1, tail[1]}
		}
	} else {
		// diagonol from eachother
		xDiff := head[0] - tail[0]
		yDiff := head[1] - tail[1]
		newTail := []int{tail[0], tail[1]}
		if xDiff > 0 {
			newTail[0] += 1
		} else {
			newTail[0] -= 1
		}
		if yDiff > 0 {
			newTail[1] += 1
		} else {
			newTail[1] -= 1
		}
		return newTail
	}
}

func parseLine(line string) (string, int) {
	parts := strings.Split(line, " ")
	return parts[0], goutil.Atoi(parts[1])
}

func areCloseEnough(head, tail []int) bool {
	if head[0] == tail[0] || head[1] == tail[1] {
		d := goutil.Dist(head, tail)
		return d <= 1
	}
	d := goutil.Dist(head, tail)
	return d <= math.Sqrt2

}
