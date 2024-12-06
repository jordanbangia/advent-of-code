package main

import (
	"fmt"
	"sort"
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

func parseInput(inputText []string) (map[string][]string, [][]string) {
	parseRules := true

	rules := map[string][]string{}

	updates := [][]string{}
	for _, line := range inputText {
		if line == "" {
			parseRules = false
		} else if parseRules {
			rule := strings.Split(line, "|")
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	return rules, updates
}

func p1(inputText []string) (int, error) {
	rules, updates := parseInput(inputText)

	checkUpdate := func(update []string) error {
		pageOrder := map[string]int{}
		for i, page := range update {
			pageOrder[page] = i
		}

		for i, page := range update {
			rulesForPage := rules[page]
			for _, otherPage := range rulesForPage {
				if place, exists := pageOrder[otherPage]; exists && place < i {
					return fmt.Errorf("failed page rule %s|%s", page, otherPage)
				}
			}
		}
		return nil
	}

	middleNumSum := 0
	for _, update := range updates {
		// fmt.Printf("checking %d\n", i)
		if err := checkUpdate(update); err != nil {
			// fmt.Println(err.Error())
		} else {
			// fmt.Printf("adding %s\n", update[len(update)/2])
			middleNumSum += goutil.Atoi(update[len(update)/2])
		}
	}

	return middleNumSum, nil
}

func p2(inputText []string) (int, error) {
	rules, updates := parseInput(inputText)

	checkUpdate := func(update []string) error {
		pageOrder := map[string]int{}
		for i, page := range update {
			pageOrder[page] = i
		}

		for i, page := range update {
			rulesForPage := rules[page]
			for _, otherPage := range rulesForPage {
				if place, exists := pageOrder[otherPage]; exists && place < i {
					return fmt.Errorf("failed page rule %s|%s", page, otherPage)
				}
			}
		}
		return nil
	}

	middleNumSum := 0
	for _, update := range updates {
		// fmt.Printf("checking %d\n", i)
		if err := checkUpdate(update); err != nil {
			// this is bad update, need to correct
			// we can sort the list given the rules
			// as the rules will provide our ordering
			sort.Slice(update, func(i, j int) bool {
				a := update[i]
				b := update[j]

				aRules, exist := rules[a]
				if exist {
					for _, d := range aRules {
						if d == b {
							return false
						}
					}
				}
				return true

			})
			middleNumSum += goutil.Atoi(update[len(update)/2])
		}
	}

	return middleNumSum, nil
}
