package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.SolutionMain(goutil.NewSolution(sol1, sol2))
}

func sol1(text []string) (int, error) {
	ids := strings.Split(text[0], ",")

	sumInvalidIDsInRange := func(start, end int) []int {
		fmt.Println("Checking", start, "-", end)
		invalidIDs := []int{}
		num := start
		for {
			if num > end {
				break
			}
			numStr := strconv.Itoa(num)
			n := len(numStr)
			if n%2 == 0 {
				firstHalf := numStr[:n/2]
				secondHalf := numStr[n/2:]
				if firstHalf == secondHalf {
					invalidIDs = append(invalidIDs, num)
				}
			}
			num += 1
		}
		fmt.Println("found", invalidIDs)
		return invalidIDs
	}

	invalidIDSum := 0
	for _, idRange := range ids {
		parts := strings.Split(idRange, "-")
		invalidIDs := sumInvalidIDsInRange(goutil.Atoi(parts[0]), goutil.Atoi(parts[1]))
		invalidIDSum += goutil.Sum(invalidIDs)
	}

	return invalidIDSum, nil
}

func sol2(text []string) (int, error) {
	ids := strings.Split(text[0], ",")

	checkOnePrefixIsInvalid := func(numStr, prefix string, i int) bool {
		for k := i; k < len(numStr); k += len(prefix) {
			next := numStr[k : k+len(prefix)]
			if prefix != next {
				return false
			}
		}
		return true
	}

	checkIfNumIsInvalid := func(num int) bool {
		numStr := strconv.Itoa(num)
		for i := 1; i <= len(numStr)/2; i++ {
			prefix := numStr[:i]
			if len(numStr)%len(prefix) != 0 {
				// can't possibly work because the original
				// string isn't divisible by the prefix
				// so skip this prefix
				continue
			}

			if checkOnePrefixIsInvalid(numStr, prefix, i) {
				return true
			}
		}
		return false
	}

	sumInvalidIDsInRange := func(start, end int) []int {
		fmt.Println("Checking", start, "-", end)
		invalidIDs := []int{}
		num := start
		for {
			if num > end {
				break
			}
			if checkIfNumIsInvalid(num) {
				invalidIDs = append(invalidIDs, num)
			}
			num += 1
		}
		fmt.Println("found", invalidIDs)
		return invalidIDs
	}

	invalidIDSum := 0
	for _, idRange := range ids {
		parts := strings.Split(idRange, "-")
		invalidIDs := sumInvalidIDsInRange(goutil.Atoi(parts[0]), goutil.Atoi(parts[1]))
		invalidIDSum += goutil.Sum(invalidIDs)
	}

	return invalidIDSum, nil
}
