package main

import (
	"fmt"
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

type line struct {
	target int
	vals   []int
}

func concat(a, b int) int {
	return goutil.Atoi(fmt.Sprintf("%d%d", a, b))
}

func solveEquation(target int, vals []int, includeConcat bool) bool {
	if len(vals) <= 1 {
		panic("not enough values")
	}

	a := vals[0]
	b := vals[1]
	if len(vals) == 2 {
		result := a+b == target || a*b == target
		if includeConcat {
			result = result || concat(a, b) == target
		}
		return result
	}

	// we have more than 2 values, so try evaluating down 2 paths
	// where we compute both a*b or a+b as the new first value
	// + the rest of the list
	// we can use an early stop condition though, since all of the vals
	// are postive and we're only doing addition and multiplication
	// if at any point a+b or a*b is greater than target, we can terminate that
	// branch as it will never get any lower.

	copiedVals := make([]int, len(vals)-1)
	copy(copiedVals[1:], vals[2:])

	p := a + b
	copiedVals[0] = p
	pSolves := p <= target && solveEquation(target, copiedVals, includeConcat)

	m := a * b
	copiedVals[0] = m
	mSolves := m <= target && solveEquation(target, copiedVals, includeConcat)

	c := concat(a, b)
	copiedVals[0] = c
	cSolves := includeConcat && c <= target && solveEquation(target, copiedVals, includeConcat)

	return pSolves || mSolves || cSolves
}

func parseInput(inputText []string) []*line {
	lines := make([]*line, len(inputText))

	for i, l := range inputText {
		line := &line{}

		parts := strings.Split(l, ": ")
		line.target = goutil.Atoi(parts[0])

		vals := strings.Split(parts[1], " ")
		valNums := []int{}
		for _, v := range vals {
			if v != "" {
				valNums = append(valNums, goutil.Atoi(v))
			}
		}
		line.vals = valNums
		lines[i] = line
	}
	return lines
}

func p1(inputText []string) (int, error) {
	sumValues := 0
	for _, line := range parseInput(inputText) {
		if solveEquation(line.target, line.vals, false) {
			sumValues += line.target
		}
	}
	return sumValues, nil
}

func p2(inputText []string) (int, error) {
	sumValues := 0
	for _, line := range parseInput(inputText) {
		if solveEquation(line.target, line.vals, true) {
			sumValues += line.target
		}
	}
	return sumValues, nil
}
