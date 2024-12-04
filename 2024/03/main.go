package main

import (
	"fmt"
	"regexp"
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
	// match something like mul(3 digits,3 digits), and capture the value
	// of the two 3 digits numbers in groups
	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	result := 0
	for _, line := range inputText {
		matches := reg.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}

		fmt.Printf("%d matches\n", len(matches))

		for _, match := range matches {
			num1 := goutil.Atoi(match[1])
			num2 := goutil.Atoi(match[2])
			// fmt.Printf("%s = %d * %d = %d\n", match[0], num1, num2, num1*num2)
			result += (num1 * num2)
		}
	}

	return result, nil
}

func p2(inputText []string) (int, error) {
	// same as above, but also include captures for do() / don't()
	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	enabled := true
	result := 0
	for _, line := range inputText {
		matches := reg.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}

		for _, match := range matches {
			if strings.Contains(match[0], "mul") {
				if enabled {
					n := (goutil.Atoi(match[1]) * goutil.Atoi(match[2]))
					fmt.Println("adding", n)
					result += n
				}
			} else if match[0] == "do()" {
				enabled = true
				fmt.Println("enabling")
			} else if match[0] == "don't()" {
				enabled = false
				fmt.Println("disabling")
			}
		}
	}

	return result, nil
}
