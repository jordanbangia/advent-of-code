package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			nil, p2,
		),
		false,
	)
}

func p1(inputText []string) (int, error) {
	reg := regexp.MustCompile(`(-?\d*),(-?\d*)`)

	seconds := 100

	maxX := 101
	maxY := 103

	quadrantCount := map[int]int{}
	initialPoints := map[string]int{}
	finalPoints := map[string]int{}

	quadrantBounds := [][][]int{
		{{-1, -1}, {(maxX - 1) / 2, (maxY - 1) / 2}},
		{{-1, (maxY - 1) / 2}, {(maxX - 1) / 2, maxY}},
		{{(maxX - 1) / 2, -1}, {maxX, (maxY - 1) / 2}},
		{{(maxX - 1) / 2, (maxY - 1) / 2}, {maxX, maxY}},
	}

	for _, line := range inputText {
		matches := reg.FindAllStringSubmatch(line, -1)

		p := []int{goutil.Atoi(matches[0][1]), goutil.Atoi(matches[0][2])}
		v := []int{goutil.Atoi(matches[1][1]), goutil.Atoi(matches[1][2])}

		initialPoints[goutil.AKey(p)] += 1

		finalP := []int{
			(p[0] + seconds*v[0]) % maxX,
			(p[1] + seconds*v[1]) % maxY,
		}

		if finalP[0] < 0 {
			finalP[0] += maxX
		}
		if finalP[1] < 0 {
			finalP[1] += maxY
		}

		finalPoints[goutil.AKey(finalP)] += 1

		for i, q := range quadrantBounds {
			min := q[0]
			max := q[1]

			xInBounds := min[0] < finalP[0] && finalP[0] < max[0]
			yInBounds := min[1] < finalP[1] && finalP[1] < max[1]
			if xInBounds && yInBounds {
				quadrantCount[i] += 1
				break
			}
		}
	}

	safetyFactor := 1
	for _, c := range quadrantCount {
		safetyFactor *= c
	}

	// fmt.Println(finalPoints)

	// fmt.Println("Initial Board:")
	// printBoard(initialPoints, maxX, maxY)

	fmt.Println("Final Board:")
	printBoard(finalPoints, maxX, maxY)

	return safetyFactor, nil
}

type Robot struct {
	p, v []int
}

func p2(inputText []string) (int, error) {
	reg := regexp.MustCompile(`(-?\d*),(-?\d*)`)

	maxX := 101
	maxY := 103

	robots := []*Robot{}
	for _, line := range inputText {
		matches := reg.FindAllStringSubmatch(line, -1)

		p := []int{goutil.Atoi(matches[0][1]), goutil.Atoi(matches[0][2])}
		v := []int{goutil.Atoi(matches[1][1]), goutil.Atoi(matches[1][2])}

		robots = append(robots, &Robot{p: p, v: v})
	}

	checkForLine := func(ix, iy int, points map[string]int) bool {
		lineLength := 0
		x := ix
		y := iy

		// check down to the right
		for {
			x += 1
			y += 1
			if _, exists := points[goutil.Key(x, y)]; !exists {
				break
			}
			lineLength += 1
		}

		// and up to the left
		x = ix
		y = iy
		for {
			x -= 1
			y -= 1
			if _, exists := points[goutil.Key(x, y)]; !exists {
				break
			}
			lineLength += 1
		}

		// very stupidly this kind of works
		// its good enough to detect a couple of false positives
		// and we just need to figure out one point where we hit it
		if lineLength > 5 {
			return true
		}

		lineLength = 0

		// instead check down to the left
		x = ix
		y = iy
		for {
			x -= 1
			y += 1
			if _, exists := points[goutil.Key(x, y)]; !exists {
				break
			}
			lineLength += 1
		}
		// and up to the right
		x = ix
		y = iy
		for {
			x += 1
			y -= 1
			if _, exists := points[goutil.Key(x, y)]; !exists {
				break
			}
			lineLength += 1
		}

		if lineLength > 5 {
			return true
		}
		return false
	}

	seconds := 0
	for {
		seconds += 1

		points := map[string]int{}
		for _, r := range robots {
			finalP := []int{
				(r.p[0] + seconds*r.v[0]) % maxX,
				(r.p[1] + seconds*r.v[1]) % maxY,
			}
			if finalP[0] < 0 {
				finalP[0] += maxX
			}
			if finalP[1] < 0 {
				finalP[1] += maxY
			}
			points[goutil.AKey(finalP)] += 1
		}

		// if this is going to be a christmas tree
		// a good number of points are going to have to exist in a straight line
		// pick some random point

		for p := range points {
			x, y := goutil.SplitKey(p)

			if checkForLine(x, y, points) {
				fmt.Println(seconds)
				printBoard(points, maxX, maxY)
				break
			}
		}

		if seconds == 10000 {
			fmt.Println("Stopping after:", seconds)
			printBoard(points, maxX, maxY)
			break
		}
	}

	return -1, nil
}

func printBoard(locs map[string]int, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		sb := strings.Builder{}
		for x := 0; x < maxX; x++ {
			c := locs[goutil.Key(x, y)]
			if c == 0 {
				sb.WriteRune('.')
			} else {
				sb.WriteString(strconv.Itoa(c))
			}
		}
		fmt.Println(sb.String())
	}
}
