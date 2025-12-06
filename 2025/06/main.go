package main

import (
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.SolutionMain(goutil.NewSolution(part1, part2))
}

type Problem struct {
	nums []int
	op   string
}

func (p *Problem) eval() int {
	v := 0
	if p.op == "*" {
		v = 1
	}

	for _, n := range p.nums {
		if n == 0 {
			continue
		}
		switch p.op {
		case "+":
			v += n
		case "*":
			v *= n
		}
	}
	return v
}

func part1(inputText []string) (int, error) {
	problems := []*Problem{}

	for i, l := range inputText {
		for k, n := range strings.Fields(l) {
			if i == 0 {
				problems = append(problems, &Problem{nums: []int{}})
			}
			if i == len(inputText)-1 {
				problems[k].op = n
			} else {
				problems[k].nums = append(problems[k].nums, goutil.Atoi(n))
			}
		}
	}

	total := 0
	for _, p := range problems {
		total += p.eval()
	}
	return total, nil
}

func part2(inputText []string) (int, error) {
	// effectively trying to transpose

	sheetWidth := 0
	for _, l := range inputText {
		sheetWidth = max(sheetWidth, len(l))
	}

	sheet := make([][]rune, sheetWidth)
	for i := range sheet {
		sheet[i] = make([]rune, len(inputText))
	}

	for i, l := range inputText {
		for k, c := range l {
			sheet[k][i] = c
		}
	}

	total := 0

	problem := &Problem{nums: []int{}}
	for i := len(sheet) - 1; i >= 0; i-- {
		numBuilder := strings.Builder{}
		for _, c := range sheet[i] {
			switch c {
			case ' ':
			case 0:
			case '+':
				problem.op = "+"
			case '*':
				problem.op = "*"
			default:
				numBuilder.WriteRune(c)
			}
		}
		problem.nums = append(problem.nums, goutil.Atoi(numBuilder.String()))
		if problem.op != "" {
			total += problem.eval()
			problem = &Problem{nums: []int{}}
		}
	}
	return total, nil
}
