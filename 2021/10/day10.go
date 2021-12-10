package main

import (
	"fmt"
	"sort"

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
	score := 0
	for _, l := range input {
		score += checkLineForIncorrectSyntax(l)
	}

	fmt.Printf("Solution 1: %d\n", score)
	return nil
}

func checkLineForIncorrectSyntax(l string) int {
	s := []rune{}

	var x rune
	for _, i := range l {
		if i == '(' || i == '[' || i == '<' || i == '{' {
			s = append(s, i)
		} else if i == ')' {
			x, s = s[len(s)-1], s[:len(s)-1]
			if x != '(' {
				// syntax error
				return 3
			}
		} else if i == ']' {
			x, s = s[len(s)-1], s[:len(s)-1]
			if x != '[' {
				return 57
			}
		} else if i == '}' {
			x, s = s[len(s)-1], s[:len(s)-1]
			if x != '{' {
				return 1197
			}
		} else if i == '>' {
			x, s = s[len(s)-1], s[:len(s)-1]
			if x != '<' {
				return 25137
			}
		}
	}

	return 0
}

func solve2(input []string) error {
	scores := []int{}
	for _, l := range input {
		if checkLineForIncorrectSyntax(l) != 0 {
			continue
		}
		scores = append(scores, completeLine(l))
	}

	sort.Ints(scores)

	fmt.Printf("Solution 2: %d\n", scores[len(scores)/2])
	return nil
}

func completeLine(l string) int {
	s := []rune{}

	for _, i := range l {
		if i == '(' || i == '[' || i == '<' || i == '{' {
			// append to stack on an open
			s = append(s, i)
		} else if i == ')' || i == ']' || i == '>' || i == '}' {
			// pop from stack on the close
			_, s = s[len(s)-1], s[:len(s)-1]
		}
	}

	score := 0
	// our stack is unfinshed

	for len(s) > 0 {
		var x rune
		x, s = s[len(s)-1], s[:len(s)-1]

		score = score * 5

		switch x {
		case '[':
			score += 2
		case '(':
			score += 1
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}

	return score
}
