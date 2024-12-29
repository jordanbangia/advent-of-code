package main

import (
	"fmt"
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

type E struct {
	x, y  int
	steps int
}

var steps = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func p1(inputText []string) (int, error) {
	maxRow := 70
	maxCol := 70
	bytesToUse := 1024

	g := goutil.Grid[string]{
		G:      map[string]string{},
		MaxRow: maxRow,
		MaxCol: maxCol,
	}

	for i, line := range inputText {
		if i < bytesToUse {
			g.G[strings.Replace(line, ",", "_", -1)] = "x"
		}
	}

	// f, _ := os.Create("/home/jordan/code/advent-of-code/2024/18/grid.txt")
	// for y := 0; y < 71; y++ {
	// 	s := strings.Builder{}
	// 	for x := 0; x < 71; x++ {
	// 		if g.G[goutil.Key(x, y)] == "x" {
	// 			s.WriteString("x")
	// 		} else {
	// 			s.WriteString(".")
	// 		}
	// 	}
	// 	s.WriteString("\n")
	// 	f.WriteString(s.String())
	// }
	// f.Close()

	target := goutil.Key(maxRow, maxCol)
	return findPath(g, target)
}

func p2(inputText []string) (int, error) {
	maxRow := 70
	maxCol := 70

	g := goutil.Grid[string]{
		G:      map[string]string{},
		MaxRow: maxRow,
		MaxCol: maxCol,
	}
	target := goutil.Key(maxRow, maxCol)
	for _, line := range inputText {
		g.G[strings.Replace(line, ",", "_", -1)] = "x"

		steps, _ := findPath(g, target)
		if steps == -1 {
			fmt.Println(line)
			return -1, nil
		}
	}
	return -1, nil
}

func findPath(g goutil.Grid[string], target string) (int, error) {
	stack := goutil.NewQueue[*E]()
	stack.Push(&E{x: 0, y: 0, steps: 0})

	visited := map[string]int{}

	for !stack.IsEmpty() {
		c, _ := stack.Pop()

		if _, alreadyVisited := visited[goutil.Key(c.x, c.y)]; alreadyVisited {
			// we've already visited this spot with a lower score
			// so there's no point coming back with a higher score
			continue

		}
		visited[goutil.Key(c.x, c.y)] = c.steps

		for _, step := range steps {
			nX := c.x + step[0]
			nY := c.y + step[1]

			if nX < 0 || nY < 0 || nX > g.MaxCol || nY > g.MaxRow {
				continue
			}

			if _, exists := g.G[goutil.Key(nX, nY)]; exists {
				// this is a corrupt space
				continue
			}

			if goutil.Key(nX, nY) == target {
				return c.steps + 1, nil
			}

			stack.Push(&E{
				x: nX, y: nY, steps: c.steps + 1,
			})
		}
	}
	return -1, nil
}
