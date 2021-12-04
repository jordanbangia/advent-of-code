package main

import (
	"fmt"
	"strconv"

	"github.com/jordanbangia/advent-of-code/goutil"
)

var runeMap map[int]byte = map[int]byte{
	1: '1',
	0: '0',
}

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
	maxOnes := len(input)
	half := maxOnes / 2

	valLength := len(input[0])

	posOneCounts := map[int]int{}

	// count the number of ones in each position
	for _, l := range input {
		for pos, car := range l {
			if car == '1' {
				posOneCounts[pos] = posOneCounts[pos] + 1
			}
		}
	}

	gamma := ""
	epsilon := ""

	for i := 0; i < valLength; i++ {
		if posOneCounts[i] > half {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Solution 1: %d\n", gammaInt*epsilonInt)

	return nil
}

func solve2(input []string) error {
	ogrInt := findOGR(input)

	csrInt := findCSR(input)

	fmt.Printf("Solution 2: %d\n", ogrInt*csrInt)

	return nil
}

func findOGR(input []string) int {
	bitPosition := 0
	filtered := input
	ogrFound := false
	ogr := ""
	for !ogrFound {
		bit := mostCommonBitAtPosition(filtered, bitPosition)

		next := []string{}
		for _, l := range filtered {
			if l[bitPosition] == runeMap[bit] {
				next = append(next, l)
			}
		}

		if len(next) == 1 {
			ogr = next[0]
			ogrFound = true
		} else {
			filtered = next
			bitPosition++
		}
	}
	ogrInt, _ := strconv.ParseInt(ogr, 2, 64)
	return int(ogrInt)
}

func findCSR(input []string) int {
	bitPosition := 0
	filtered := input
	csrFound := false
	csr := ""
	for !csrFound {
		bit := mostCommonBitAtPosition(filtered, bitPosition)
		// flip it cause we actually want the least common
		if bit == 1 {
			bit = 0
		} else {
			bit = 1
		}

		next := []string{}
		for _, l := range filtered {
			if l[bitPosition] == runeMap[bit] {
				next = append(next, l)
			}
		}

		if len(next) == 1 {
			csr = next[0]
			csrFound = true
		} else {
			filtered = next
			bitPosition++
		}
	}

	csrInt, _ := strconv.ParseInt(csr, 2, 64)
	return int(csrInt)
}

func mostCommonBitAtPosition(input []string, position int) int {
	oneCount := 0
	zeroCount := 0
	for _, l := range input {
		if l[position] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}

	if oneCount >= zeroCount {
		return 1
	}
	return 0
}
