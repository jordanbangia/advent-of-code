package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = solve1(input)
	_ = solve2(input)

}

type Board struct {
	b     [][]int
	eles  map[int][]int
	marks [][]bool
}

func NewBoard(i []string) *Board {
	b := [][]int{}

	for _, l := range i {
		if l == "" {
			continue
		}
		row := []int{}
		for _, c := range strings.Split(l, " ") {
			if c == "" {
				continue
			}
			n, _ := strconv.Atoi(c)
			row = append(row, n)
		}
		b = append(b, row)
	}

	m := make([][]bool, 5)
	for i := 0; i < 5; i++ {
		m[i] = []bool{false, false, false, false, false}
	}

	eles := map[int][]int{}
	for i, row := range b {
		for j, v := range row {
			eles[v] = []int{i, j}
		}
	}

	return &Board{b: b, eles: eles, marks: m}
}

func (b *Board) callNums(nums []int) (int, int) {
	for i, n := range nums {
		if coord, ok := b.eles[n]; ok {
			b.marks[coord[0]][coord[1]] = true
			if b.isWin(coord[0], coord[1]) {
				return i, b.sumUnmarked() * n
			}
		}
		// fmt.Println(b.marks)
	}
	return len(nums) + 100, -1
}

func (b *Board) isWin(i, j int) bool {
	iWin := true
	for x := 0; x < 5; x++ {
		iWin = iWin && b.marks[i][x]
	}

	jWin := true
	for x := 0; x < 5; x++ {
		jWin = jWin && b.marks[x][j]
	}

	return iWin || jWin
}

func (b *Board) sumUnmarked() int {
	s := 0
	for i, row := range b.marks {
		for j, marked := range row {
			if !marked {
				s += b.b[i][j]
			}
		}
	}
	return s
}

func solve1(input []string) error {
	// line 0 is numbers called
	calledNumbers := []int{}
	for _, c := range strings.Split(input[0], ",") {
		n, _ := strconv.Atoi(c)
		calledNumbers = append(calledNumbers, n)
	}
	// fmt.Println(calledNumbers)

	fastestWin := len(calledNumbers) + 100
	fastestWinScore := -1

	for i := 2; i < len(input); i += 6 {
		b := NewBoard(input[i : i+5])
		// fmt.Println(b)

		turnsToWin, score := b.callNums(calledNumbers)
		if turnsToWin < fastestWin {
			fastestWin = turnsToWin
			fastestWinScore = score
		}
	}

	fmt.Printf("Solution 1: %d turns, %d score\n", fastestWin, fastestWinScore)

	return nil
}

func solve2(input []string) error {
	// line 0 is numbers called
	calledNumbers := []int{}
	for _, c := range strings.Split(input[0], ",") {
		n, _ := strconv.Atoi(c)
		calledNumbers = append(calledNumbers, n)
	}
	// fmt.Println(calledNumbers)

	slowestWin := 0
	slowestWinScore := -1

	for i := 2; i < len(input); i += 6 {
		b := NewBoard(input[i : i+5])
		// fmt.Println(b)

		turnsToWin, score := b.callNums(calledNumbers)
		if turnsToWin > slowestWin {
			slowestWin = turnsToWin
			slowestWinScore = score
		}
	}

	fmt.Printf("Solution 2: %d turns, %d score\n", slowestWin, slowestWinScore)

	return nil
}
