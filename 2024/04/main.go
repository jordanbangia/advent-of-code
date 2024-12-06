package main

import (
	"errors"

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

func p1(inputText []string) (int, error) {
	coords := []goutil.Coord{}

	letters := map[string]rune{}

	checkForXmas := func(coord goutil.Coord, dirX, dirY int) error {
		curX, curY := coord.X, coord.Y
		checkWord := []rune{'X', 'M', 'A', 'S'}
		i := 0
		for {
			if i == len(checkWord) {
				return nil
			}

			if letters[goutil.Key(curX, curY)] != checkWord[i] {
				return errors.New("not a complete word")
			}

			i += 1
			curX += dirX
			curY += dirY
		}
	}

	// parse the input into a grid
	// also figure out which coords are x's
	// thats where we can start our search from
	for i, line := range inputText {
		for k, c := range line {
			letters[goutil.Key(i, k)] = c
			if c == 'X' {
				coords = append(coords, goutil.Coord{X: i, Y: k})
			}
		}
	}

	completeWords := 0

	for _, coord := range coords {
		directions := [][]int{
			{0, 1},   // forward --> positive Y
			{0, -1},  // backwards --> negative Y
			{1, 0},   // down --> postive X
			{-1, 0},  // up --> postive Y
			{1, 1},   // diagonal down and right
			{1, -1},  // diagonal down and left
			{-1, 1},  // diagonal up and right
			{-1, -1}, // diagonal up and left
		}

		for _, dir := range directions {
			if err := checkForXmas(coord, dir[0], dir[1]); err == nil {
				completeWords += 1
			}
		}
	}
	return completeWords, nil
}

func p2(inputText []string) (int, error) {
	coords := []goutil.Coord{}

	letters := map[string]rune{}

	// this time we pull the A's because we need to find shapes like
	for i, line := range inputText {
		for k, c := range line {
			letters[goutil.Key(i, k)] = c
			if c == 'A' {
				coords = append(coords, goutil.Coord{X: i, Y: k})
			}
		}
	}

	checkForXMAS := func(coord goutil.Coord) error {
		if letters[coord.Key()] != 'A' {
			return errors.New("coord isn't at A")
		}

		topL, bottomR := letters[goutil.Key(coord.X-1, coord.Y-1)], letters[goutil.Key(coord.X+1, coord.Y+1)]
		topR, bottomL := letters[goutil.Key(coord.X-1, coord.Y+1)], letters[goutil.Key(coord.X+1, coord.Y-1)]

		if ((topL == 'M' && bottomR == 'S') || (topL == 'S' && bottomR == 'M')) &&
			((topR == 'M' && bottomL == 'S') || (topR == 'S' && bottomL == 'M')) {
			return nil
		}
		return errors.New("it isn't an x")
	}

	XMASSES := 0

	for _, coord := range coords {
		if err := checkForXMAS(coord); err == nil {
			XMASSES += 1
		}
	}

	return XMASSES, nil
}
