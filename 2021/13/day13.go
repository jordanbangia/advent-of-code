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

	_ = solve(input)
}

type fold struct {
	axis string
	val  int
}

func solve(input []string) error {

	coords := map[string]bool{}

	readCoordinates := true

	folds := []fold{}

	for _, l := range input {
		if l == "" {
			readCoordinates = false
			continue
		}

		if readCoordinates {
			coords[l] = true
		} else {
			l = strings.Replace(l, "fold along ", "", -1)
			parts := strings.Split(l, "=")
			f := fold{}
			f.axis = parts[0]
			val, _ := strconv.Atoi(parts[1])
			f.val = val
			folds = append(folds, f)
		}
	}

	// fmt.Println(coords)
	for i, f := range folds {
		coords = applyFold(f.axis, f.val, coords)
		fmt.Printf("After %d folds: %d\n", i+1, len(coords))
	}

	printCoords(coords)

	return nil
}

func applyFold(axis string, val int, coords map[string]bool) map[string]bool {
	next := map[string]bool{}

	for c := range coords {
		x, y := splitCoords(c)

		if axis == "y" {
			if y < val {
				next[key(x, y)] = true
			} else {
				next[key(x, val-(y-val))] = true
			}
		} else {
			if x < val {
				next[key(x, y)] = true
			} else {
				next[key(val-(x-val), y)] = true
			}
		}
	}
	return next
}

func splitCoords(s string) (int, int) {
	p := strings.Split(s, ",")

	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])

	return x, y
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printCoords(coords map[string]bool) {
	maxX := -1
	maxY := -1

	for c := range coords {
		x, y := splitCoords(c)
		maxX = max(x, maxX)
		maxY = max(y, maxY)
	}

	outs := make([]string, maxY+1)

	for i := range outs {
		outs[i] = strings.Repeat(".", maxX+1)
	}

	for c := range coords {
		x, y := splitCoords(c)

		outs[y] = outs[y][:x] + "#" + outs[y][x+1:]
	}

	for _, o := range outs {
		fmt.Println(o)
	}
}
