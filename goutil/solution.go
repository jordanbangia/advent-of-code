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
