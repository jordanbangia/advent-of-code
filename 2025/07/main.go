package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() { goutil.SolutionMain(goutil.NewSolution(part1, part2)) }

func parseManifold(inputText []string) [][]string {
	manifold := make([][]string, len(inputText))

	for i, l := range inputText {
		elems := make([]string, len(l))
		for k, c := range l {
			elems[k] = string(c)
		}
		manifold[i] = elems
	}
	return manifold
}

func part1(inputText []string) (int, error) {
	manifold := parseManifold(inputText)
	splitsHit := 0
	for col, line := range manifold {
		if col == len(manifold)-1 {
			break
		}
		for i, x := range line {
			switch x {
			case "S":
				fallthrough
			case "|":
				switch below := manifold[col+1][i]; below {
				case ".":
					fallthrough
				case "|":
					manifold[col+1][i] = "|"
				case "^":
					splitsHit += 1
					// we're going to fall onto a splitter, so split
					manifold[col+1][i-1] = "|"
					manifold[col+1][i+1] = "|"
				}
			}
		}
	}

	return splitsHit, nil
}

type State struct {
	row  int
	col  int
	path string
}

func part2(inputText []string) (int, error) {
	manifold := parseManifold(inputText)

	start := func() int {
		for r, x := range manifold[0] {
			if x == "S" {
				return r
			}
		}
		return -1
	}()

	beams := make([][]int, len(manifold))
	for i := range manifold {
		beams[i] = make([]int, len(manifold[i]))
	}

	beams[0][start] = 1

	// should print the grid / beams after each iteration
	printBeams := func() {
		sb := strings.Builder{}
		for _, l := range manifold {
			for _, c := range l {
				// if beams[x][y] != 0 {
				// 	sb.WriteString(fmt.Sprintf("%d", (beams[x][y])))
				// } else {
				sb.WriteString(c)
				// }
			}
			sb.WriteRune('\n')
		}
		fmt.Println(sb.String())
	}

	for col, line := range manifold {
		if col == len(manifold)-1 {
			break
		}
		for i, x := range line {
			switch x {
			case "S":
				fallthrough
			case "|":
				switch below := manifold[col+1][i]; below {
				case ".":
					fallthrough
				case "|":
					manifold[col+1][i] = "|"
					beams[col+1][i] += beams[col][i]
				case "^":
					manifold[col+1][i-1] = "|"
					manifold[col+1][i+1] = "|"
					beams[col+1][i-1] += beams[col][i]
					beams[col+1][i+1] += beams[col][i]
				}
			}
		}
	}

	printBeams()
	endStates := goutil.Sum(beams[len(beams)-1])
	return endStates, nil
}
