package main

import (
	"fmt"

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

func nodesAreEqual(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func p1(inputText []string) (int, error) {
	/*
		We define an Antinode at position X if there exists
		2 nodes on a line joining X and d(X, n1) = d(x, n2) / 2
		(i.e. n2 is twice as far from X as n1)

		Therefore we need to consider checking every possible square.

		But, for every square, we would need to check every possible line.
		Instead, what if we checked every possible pair of nodes of the
		same type.

		For every pair of nodes, compute the line (y=mx+b) that goes through
		those 2 nodes.  Determine the 2 possible antinodes that could exist
		on those lines, given the distance issue above.
		If that works out, we can consider that point an antinode, add it to
		our unique list.
	*/

	antennas := map[string][][]int{}

	maxRow := len(inputText)
	maxCol := -1

	for i, line := range inputText {
		maxCol = len(line)
		for j, l := range line {
			if l != '.' {
				antennas[string(l)] = append(antennas[string(l)], []int{i, j})
			}
		}
	}

	nodeIsValid := func(a []int) bool {
		return a[0] >= 0 && a[0] < maxRow && a[1] >= 0 && a[1] < maxCol
	}

	antenaLocations := map[string]struct{}{}

	for receiverName, receiverPositions := range antennas {
		fmt.Printf("Working on %s\n", receiverName)

		for i := 0; i < len(receiverPositions); i++ {
			for j := i + 1; j < len(receiverPositions); j++ {

				n1 := receiverPositions[i]
				n2 := receiverPositions[j]

				rowDiff := n2[0] - n1[0]
				colDiff := n2[1] - n1[1]

				// given the diffs, add / sub them from each point
				// one will overlap the other so can ignore that one

				possibleAntinode := []int{n1[0] + rowDiff, n1[1] + colDiff}
				if !nodesAreEqual(possibleAntinode, n1) && !nodesAreEqual(possibleAntinode, n2) && nodeIsValid(possibleAntinode) {
					antenaLocations[goutil.Key(possibleAntinode[0], possibleAntinode[1])] = struct{}{}
				}

				possibleAntinode = []int{n1[0] - rowDiff, n1[1] - colDiff}
				if !nodesAreEqual(possibleAntinode, n1) && !nodesAreEqual(possibleAntinode, n2) && nodeIsValid(possibleAntinode) {
					antenaLocations[goutil.Key(possibleAntinode[0], possibleAntinode[1])] = struct{}{}
				}

				possibleAntinode = []int{n2[0] + rowDiff, n2[1] + colDiff}
				if !nodesAreEqual(possibleAntinode, n1) && !nodesAreEqual(possibleAntinode, n2) && nodeIsValid(possibleAntinode) {
					antenaLocations[goutil.Key(possibleAntinode[0], possibleAntinode[1])] = struct{}{}
				}

				possibleAntinode = []int{n2[0] - rowDiff, n2[1] - colDiff}
				if !nodesAreEqual(possibleAntinode, n1) && !nodesAreEqual(possibleAntinode, n2) && nodeIsValid(possibleAntinode) {
					antenaLocations[goutil.Key(possibleAntinode[0], possibleAntinode[1])] = struct{}{}
				}
			}
		}
	}

	return len(antenaLocations), nil
}

func p2(inputText []string) (int, error) {

	antennas := map[string][][]int{}

	maxRow := len(inputText)
	maxCol := -1

	for i, line := range inputText {
		maxCol = len(line)
		for j, l := range line {
			if l != '.' {
				antennas[string(l)] = append(antennas[string(l)], []int{i, j})
			}
		}
	}

	nodeIsValid := func(a []int) bool {
		return a[0] >= 0 && a[0] < maxRow && a[1] >= 0 && a[1] < maxCol
	}

	antenaLocations := map[string]struct{}{}

	for receiverName, receiverPositions := range antennas {
		fmt.Printf("Working on %s\n", receiverName)

		for i := 0; i < len(receiverPositions); i++ {
			for j := i + 1; j < len(receiverPositions); j++ {

				n1 := receiverPositions[i]
				n2 := receiverPositions[j]

				rowDiff := n2[0] - n1[0]
				colDiff := n2[1] - n1[1]

				// given the diffs, add / sub them from each point
				// one will overlap the other so can ignore that one

				// go one way from n1
				possibleAntinode := []int{n1[0], n1[1]}
				for {
					if !nodeIsValid(possibleAntinode) {
						break
					}
					antenaLocations[goutil.Key(possibleAntinode[0], possibleAntinode[1])] = struct{}{}
					possibleAntinode[0] += rowDiff
					possibleAntinode[1] += colDiff
				}

				// go the other direction from n1
				possibleAntinode = []int{n1[0], n1[1]}
				for {
					if !nodeIsValid(possibleAntinode) {
						break
					}
					antenaLocations[goutil.Key(possibleAntinode[0], possibleAntinode[1])] = struct{}{}
					possibleAntinode[0] -= rowDiff
					possibleAntinode[1] -= colDiff
				}
			}
		}
	}

	return len(antenaLocations), nil
}
