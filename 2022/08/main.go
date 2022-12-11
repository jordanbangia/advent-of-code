package main

import (
	"fmt"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := parseInput(input)
	fmt.Println(countVisible(grid))
	fmt.Println(scenicScore(grid))
}

func parseInput(input []string) [][]int {
	grid := [][]int{}
	for _, line := range input {
		row := []int{}
		for _, c := range line {
			row = append(row, goutil.Atoi(string(c)))
		}
		grid = append(grid, row)
	}
	return grid
}

func countVisible(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	isVisible := map[string]bool{}

	for i := 0; i < rows; i++ {
		maxSeen := grid[i][0]
		for j := 0; j < cols; j++ {
			if j == 0 {
				isVisible[key(i, j)] = true
			} else if grid[i][j] > maxSeen {
				isVisible[key(i, j)] = true
				maxSeen = grid[i][j]
			}
		}

		maxSeen = grid[i][cols-1]
		for j := cols - 1; j >= 0; j-- {
			if j == cols-1 {
				isVisible[key(i, j)] = true
			} else if grid[i][j] > maxSeen {
				isVisible[key(i, j)] = true
				maxSeen = grid[i][j]
			}
		}
	}

	for j := 0; j < cols; j++ {
		maxSeen := grid[0][j]
		for i := 0; i < rows; i++ {
			if i == 0 {
				isVisible[key(i, j)] = true
			} else if grid[i][j] > maxSeen {
				isVisible[key(i, j)] = true
				maxSeen = grid[i][j]
			}
		}

		maxSeen = grid[rows-1][j]
		for i := rows - 1; i >= 0; i-- {
			if i == rows-1 {
				isVisible[key(i, j)] = true
			} else if grid[i][j] > maxSeen {
				isVisible[key(i, j)] = true
				maxSeen = grid[i][j]
			}
		}
	}

	return len(isVisible)
}

func scenicScore(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	scores := map[string]int{}

	maxScore := -1

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			v := grid[i][j]

			// go left
			leftScore := 1
			if j != 0 {
				for l := j - 1; l > 0 && v > grid[i][l]; l-- {
					leftScore += 1
				}
			} else {
				leftScore = 0
			}

			// go up
			upScore := 1
			if i != 0 {
				for l := i - 1; l > 0 && v > grid[l][j]; l-- {
					upScore += 1
				}
			} else {
				upScore = 0
			}

			// go right
			rightScore := 1
			if j != cols-1 {
				for l := j + 1; l < cols-1 && v > grid[i][l]; l++ {
					rightScore += 1
				}
			} else {
				rightScore = 0
			}

			// go down
			downScore := 1
			if i != rows-1 {
				for l := i + 1; l < rows-1 && v > grid[l][j]; l++ {
					downScore += 1
				}
			} else {
				downScore = 0
			}

			score := leftScore * rightScore * upScore * downScore
			// fmt.Println(i, j, leftScore, rightScore, upScore, downScore, score)
			scores[key(i, j)] = score
			maxScore = goutil.Max(maxScore, score)
		}
	}

	// fmt.Println(scores)
	return maxScore
}

func key(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}
