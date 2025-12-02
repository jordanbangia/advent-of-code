package goutil

import (
	"errors"
	"flag"
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

func SolutionMain(s Solution) {
	var testOnly = flag.Bool("test", false, "run test input only")
	var realOnly = flag.Bool("real", false, "run solution input only")
	var runPart1 = flag.Bool("p1", false, "run part 1 only")
	var runPart2 = flag.Bool("p2", false, "run part 2 only")
	flag.Parse()

	p1 := *runPart1 || (!*runPart1 && !*runPart2)
	p2 := *runPart2 || (!*runPart1 && !*runPart2)

	testInput := *testOnly || (!*testOnly && !*realOnly)
	realInput := *realOnly || (!*testOnly && !*realOnly)

	if testInput {
		fmt.Println("Test solutions:")
		input, _ := ReadFile("test.txt")
		if p1 {
			p1Result, _ := s.Part1(input)
			fmt.Println("Part 1: ", p1Result)
		}
		if p2 {
			p2Result, _ := s.Part2(input)
			fmt.Println("Part 2: ", p2Result)
		}
	}

	if realInput {
		fmt.Println("Puzzle solutions:")
		input, _ := ReadFile("input.txt")
		if p1 {
			p1Result, _ := s.Part1(input)
			fmt.Println("Part 1: ", p1Result)
		}
		if p2 {
			p2Result, _ := s.Part2(input)
			fmt.Println("Part 2: ", p2Result)
		}
	}

}
