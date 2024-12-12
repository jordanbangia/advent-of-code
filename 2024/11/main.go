package main

import (
	"strconv"
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

type Stone struct {
	s      int
	blinks int
}

func p1(inputText []string) (int, error) {
	return blinkAtStones(inputText, 25)
}

func p2(inputText []string) (int, error) {
	return blinkAtStones(inputText, 75)
}

func blinkAtStones(inputText []string, blinks int) (int, error) {
	line := inputText[0]
	parts := strings.Split(line, " ")

	// Instead of keeping track of every stone, we know that each stone
	// with the same engraving on it will evolve in the same way
	// Plus there will tend to be loops as certain patterns will reveal
	// and we'll likely fall back to similar numbers.
	// So we instead keep track of the number on the stone and the current
	// count of that stone.  We then apply the rule for each unique stone
	// number and do all the shifts accordingly based on the counts.

	stones := map[int]int{}
	for _, stone := range parts {
		stones[goutil.Atoi(stone)] = 1
	}

	applyRules := func(stone int) []int {
		if stone == 0 {
			return []int{1}
		}

		stoneStr := strconv.Itoa(stone)
		if len(stoneStr)%2 == 0 {
			x := []int{
				goutil.Atoi(stoneStr[:len(stoneStr)/2]),
				goutil.Atoi(stoneStr[(len(stoneStr) / 2):]),
			}
			return x
		}
		return []int{stone * 2024}
	}

	// fmt.Printf("%+v - %d\n", stones, countStones(stones))
	for i := 0; i < blinks; i++ {
		// fmt.Printf("blink %d\n", i+1)

		// blinks are calculated simultaneously for all stones
		// so keep a separate map of the updates for the given blink
		updates := map[int]int{}
		for stone, count := range stones {
			if count == 0 {
				continue
			}
			// remove the current stone, since its values will change
			updates[stone] -= count

			// calcluate what the next stones are going to be for the given
			// changing stone
			n := applyRules(stone)
			// fmt.Printf("%d, %d -> %+v\n", stone, count, n)
			for _, next := range n {
				updates[next] += count
			}

		}

		// apply the updates to the overall list of stones
		for updateStone, updateCount := range updates {
			stones[updateStone] += updateCount
			if stones[updateStone] == 0 {
				delete(stones, updateStone)
			}
		}
		// fmt.Printf("%+v - %d\n", stones, countStones(stones))
	}

	return countStones(stones), nil
}

func countStones(stones map[int]int) int {
	totalStones := 0
	for _, count := range stones {
		totalStones += count
	}
	return totalStones
}
