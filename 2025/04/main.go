package main

import "github.com/jordanbangia/advent-of-code/goutil"

func main() {
	goutil.SolutionMain(goutil.NewSolution(part1, part2))
}

var dirs = [][]int{
	{1, 1}, {1, 0}, {1, -1},
	{0, 1}, {0, -1},
	{-1, 1}, {-1, 0}, {-1, -1},
}

func findRemovableRolls(g *goutil.Grid[rune]) []string {
	accessible := func(r, c int) bool {
		rollsOfPaper := 0
		for _, d := range dirs {
			nR, nC := r+d[0], c+d[1]
			if nR < 0 || nR >= g.MaxRow {
				continue
			}
			if nC < 0 || nC >= g.MaxCol {
				continue
			}
			if g.G[goutil.Key(nR, nC)] == '@' {
				rollsOfPaper += 1
			}
		}
		return rollsOfPaper < 4
	}

	removableRolls := []string{}
	for k, v := range g.G {
		if v == '@' {
			if accessible(goutil.SplitKey(k)) {
				removableRolls = append(removableRolls, k)
			}
		}
	}
	return removableRolls
}

func buildGrid(input []string) *goutil.Grid[rune] {
	g := goutil.Grid[rune]{
		G:      map[string]rune{},
		MaxRow: len(input[0]),
		MaxCol: len(input),
	}

	for c, row := range input {
		for r, x := range row {
			g.G[goutil.Key(r, c)] = x
		}
	}
	return &g
}

func part1(input []string) (int, error) {
	return len(findRemovableRolls(buildGrid(input))), nil
}

func part2(input []string) (int, error) {
	grid := buildGrid(input)

	removedRolls := 0
	for {
		removableRolls := findRemovableRolls(grid)
		if len(removableRolls) == 0 {
			break
		}
		for _, removableRoll := range removableRolls {
			grid.G[removableRoll] = '.'
		}
		removedRolls += len(removableRolls)
	}
	return removedRolls, nil
}
