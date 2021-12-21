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

type Line struct {
	x1, y1, x2, y2 int
}

func (l Line) points() []string {
	points := []string{}
	if l.x1 == l.x2 {
		start := goutil.Min(l.y1, l.y2)
		end := goutil.Max(l.y1, l.y2)

		for i := start; i <= end; i++ {
			points = append(points, key(l.x1, i))
		}
	} else if l.y1 == l.y2 {
		start := goutil.Min(l.x1, l.x2)
		end := goutil.Max(l.x1, l.x2)

		for i := start; i <= end; i++ {
			points = append(points, key(i, l.y1))
		}
	} else {
		// its a 45 degree line -> slope = 1 or -1

		// for simplicity, swap x1, y1 to be the left most point
		x1, x2, y1, y2 := 0, 0, 0, 0
		if l.x1 < l.x2 {
			x1 = l.x1
			y1 = l.y1
			x2 = l.x2
			y2 = l.y2
		} else {
			x1 = l.x2
			y1 = l.y2
			x2 = l.x1
			y2 = l.y1
		}

		// generate all the points along the line
		m := (y2 - y1) / (x2 - x1)
		for x1 != x2 && y1 != y2 {
			points = append(points, key(x1, y1))
			x1 += 1
			y1 += m
		}
		points = append(points, key(x2, y2))
	}
	return points
}

func parseLine(l string) Line {
	parts := strings.Split(l, " -> ")

	start := parts[0]
	end := parts[1]

	startCoords := strings.Split(start, ",")
	endCoords := strings.Split(end, ",")

	x1, _ := strconv.Atoi(startCoords[0])
	y1, _ := strconv.Atoi(startCoords[1])

	x2, _ := strconv.Atoi(endCoords[0])
	y2, _ := strconv.Atoi(endCoords[1])

	return Line{x1, y1, x2, y2}
}

func countIntersections(lines []Line) int {
	coveredPoints := map[string]int{}
	for _, l := range lines {
		for _, point := range l.points() {
			coveredPoints[point] = coveredPoints[point] + 1
		}
	}

	intersections := 0
	for _, intersects := range coveredPoints {
		if intersects > 1 {
			intersections += 1
		}
	}
	return intersections
}

// we use a map to represent the 2D space, since each coordinate is unique
// and we dont' care about the empty elements
func key(a, b int) string {
	return fmt.Sprintf("%d:%d", a, b)
}

func solve1(input []string) error {
	lines := []Line{}

	for _, l := range input {
		pl := parseLine(l)

		// only consider horizontal or vertical lines
		if pl.x1 == pl.x2 || pl.y1 == pl.y2 {
			lines = append(lines, pl)
		}
	}

	fmt.Printf("Solution 1: %d intersections\n", countIntersections(lines))

	return nil
}

func solve2(input []string) error {
	lines := []Line{}

	for _, l := range input {
		pl := parseLine(l)
		lines = append(lines, pl)
	}

	fmt.Printf("Solution 2: %d intersections\n", countIntersections(lines))

	return nil
}
