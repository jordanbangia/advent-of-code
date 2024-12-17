package goutil

import (
	"errors"
	"fmt"
)

type Solution interface {
	Part1(inputText []string) (int, error)
	Part2(inputText []string) (int, error)
}

type solution struct {
	part1 func([]string) (int, error)
	part2 func([]string) (int, error)
}

func (s *solution) Part1(inputText []string) (int, error) {
	if s.part1 != nil {
		return s.part1(inputText)
	}
	return -1, errors.New("Part1 is not defined")
}

func (s *solution) Part2(inputText []string) (int, error) {
	if s.part2 != nil {
		return s.part2(inputText)
	}
	return -1, errors.New("Part2 is not defined")
}

func NewSolution(part1, part2 func(inputText []string) (int, error)) Solution {
	return &solution{
		part1: part1,
		part2: part2,
	}
}

type oneFuncSolution struct {
	f func([]string) (int, int, error)

	p1 int
	p2 int
}

func (o *oneFuncSolution) Part1(inputText []string) (int, error) {
	p1, p2, err := o.f(inputText)
	o.p2 = p2
	return p1, err
}

func (o *oneFuncSolution) Part2(inputText []string) (int, error) {
	return o.p2, nil
}

func NewOneFuncSolution(soln func(inputText []string) (int, int, error)) Solution {
	return &oneFuncSolution{
		f: soln,
	}
}

func RunSolution(s Solution, onlyTest bool) {
	input, _ := ReadFile("test.txt")

	p1, _ := s.Part1(input)
	p2, _ := s.Part2(input)

	fmt.Print("Test solutions:\n")
	PrintSolution(p1, p2)

	if onlyTest {
		return
	}

	input, _ = ReadFile("input.txt")

	p1, _ = s.Part1(input)
	p2, _ = s.Part2(input)

	fmt.Print("Puzzle solutions:\n")
	PrintSolution(p1, p2)
}
