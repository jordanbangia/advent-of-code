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

	fmt.Printf("Solution 1: %d\n", solve(input, 10))
	fmt.Println("-------")
	fmt.Printf("Solution 2: %d\n", solve(input, 40))
}

func solve(input []string, steps int) int {
	start := input[0]

	pMapping := map[string]string{}

	for _, l := range input[2:] {
		parts := strings.Split(l, " -> ")
		pMapping[parts[0]] = parts[1]
	}

	keyCache := map[key]map[rune]int{}

	n := start
	fmt.Println("Start:", n)

	elementCount := map[rune]int{}

	var previousRune rune

	for _, r := range start {
		if previousRune != 0 {
			pair := []rune{previousRune, r}
			for k, v := range applyInsertions(pair, pMapping, 0, keyCache, steps) {
				elementCount[k] += v
			}
		} else {
			elementCount[r]++
		}
		previousRune = r
	}

	mostCommonVal := -1
	leastCommonVal := -1
	for _, v := range elementCount {
		if mostCommonVal == -1 {
			mostCommonVal = v
		}
		if leastCommonVal == -1 {
			leastCommonVal = v
		}

		mostCommonVal = goutil.Max(v, mostCommonVal)
		leastCommonVal = goutil.Min(v, leastCommonVal)
	}

	return mostCommonVal - leastCommonVal
}

type key struct {
	i    int
	pair string
}

/**
the idea: we can a given pair, and extend each pair to 40 steps into the future
there is going to be some repetition at some point, and for that we use a cache
we also don't necessarily care about the string itself, just the counts as the value comes out
so cache that instead of the strings
**/
func applyInsertions(p []rune, m map[string]string, step int, mem map[key]map[rune]int, maxSteps int) map[rune]int {
	k := key{step, string(p)}
	if m, ok := mem[k]; ok {
		return m
	}

	if step == maxSteps {
		return map[rune]int{p[1]: 1}
	}
	elementCount := map[rune]int{}
	n := m[string(p)]
	for k, v := range applyInsertions([]rune{p[0], rune(n[0])}, m, step+1, mem, maxSteps) {
		elementCount[k] += v
	}
	for k, v := range applyInsertions([]rune{rune(n[0]), p[1]}, m, step+1, mem, maxSteps) {
		elementCount[k] += v
	}
	mem[k] = elementCount
	return elementCount
}

// this will error out after ~30 steps
func doStep(start string, m map[string]string) string {
	// for large strings, using a builder is much faster
	next := strings.Builder{}

	for i := 0; i < len(start)-1; i++ {
		pair := string(start[i]) + string(start[i+1])
		insert := m[pair]
		next.WriteByte(start[i])
		next.WriteString(insert)
	}
	next.WriteByte(start[len(start)-1])

	return next.String()
}

func countElements(s string) map[rune]int {
	m := map[rune]int{}

	for _, r := range s {
		m[r] += 1
	}

	return m
}
