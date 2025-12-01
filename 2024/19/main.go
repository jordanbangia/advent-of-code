package main

import (
	"slices"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			p1, p2,
		),
		false,
	)
}

func p1(inputText []string) (int, error) {

	options := strings.Split(strings.Replace(inputText[0], " ", "", -1), ",")
	slices.SortFunc(options, func(a string, b string) int {
		return len(a) - len(b)
	})
	slices.Reverse(options)

	possiblePrefixes := map[string]int{}
	for _, option := range options {
		possiblePrefixes[option] = 1
	}

	var patternCanBeMade func(string) bool

	patternCanBeMade = func(pattern string) bool {
		// we've already seen this prefix, so we can abort early
		if canBeMade := possiblePrefixes[pattern]; canBeMade == 1 {
			return true
		} else if canBeMade == -1 {
			return false
		}

		for i := len(pattern) - 1; i > -1; i-- {
			prefix := pattern[:i]
			if possiblePrefixes[prefix] == 1 {
				n := strings.Replace(pattern, prefix, "", 1)
				if n == "" {
					return true
				}
				if patternCanBeMade(n) {
					possiblePrefixes[n] = 1
					possiblePrefixes[pattern] = 1
					return true
				}
			}
		}
		possiblePrefixes[pattern] = -1
		return false
	}

	possibleDesigns := 0
	for _, pattern := range inputText[2:] {
		// fmt.Println(pattern)
		if patternCanBeMade(pattern) {
			possibleDesigns += 1
		}
	}
	return possibleDesigns, nil
}

func p2(inputText []string) (int, error) {
	options := strings.Split(strings.Replace(inputText[0], " ", "", -1), ",")

	opts := map[string]int{}

	var countWays func(string) int
	countWays = func(pattern string) (n int) {
		if n, ok := opts[pattern]; ok {
			return n
		}
		defer func() { opts[pattern] = n }()

		if pattern == "" {
			return 1
		}
		for _, s := range options {
			if strings.HasPrefix(pattern, s) {
				n += countWays(pattern[len(s):])
			}
		}
		return n
	}

	c := 0
	for _, pattern := range inputText[2:] {
		c += countWays(pattern)
	}
	return c, nil
}
