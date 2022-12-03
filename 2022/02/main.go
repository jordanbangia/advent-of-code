package main

import (
	"fmt"

	"github.com/jordanbangia/advent-of-code/goutil"
)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"

	WIN  = "win"
	LOSS = "loss"
	DRAW = "draw"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	opponentMap := map[rune]string{
		'A': ROCK,
		'B': PAPER,
		'C': SCISSORS,
	}

	responseMap := map[rune]string{
		'X': ROCK,
		'Y': PAPER,
		'Z': SCISSORS,
	}

	outcomeMap := map[rune]string{
		'X': LOSS,
		'Y': DRAW,
		'Z': WIN,
	}

	points := map[string]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}

	sol1 := 0
	sol2 := 0
	for _, line := range input {
		opponentThrow := opponentMap[rune(line[0])]
		myThrow := responseMap[rune(line[2])]

		sol1 += matchScore(opponentThrow, myThrow) + points[myThrow]

		expectedOutcome := outcomeMap[rune(line[2])]
		shouldThrow := pickThrow(opponentThrow, expectedOutcome)
		sol2 += matchScore(opponentThrow, shouldThrow) + points[shouldThrow]
	}

	goutil.PrintSolution(sol1, sol2)
}

func matchScore(opponent, me string) int {
	if opponent == me {
		// draw
		return 3
	}

	switch {
	case opponent == ROCK && me == PAPER:
		fallthrough
	case opponent == PAPER && me == SCISSORS:
		fallthrough
	case opponent == SCISSORS && me == ROCK:
		return 6
	}

	// the other 3 scenarios are all losses
	return 0
}

func pickThrow(oppent, outcome string) string {
	if outcome == DRAW {
		return oppent
	}

	switch oppent {
	case ROCK:
		if outcome == WIN {
			return PAPER
		}
		return SCISSORS
	case PAPER:
		if outcome == WIN {
			return SCISSORS
		}
		return ROCK
	case SCISSORS:
		if outcome == WIN {
			return ROCK
		}
		return PAPER
	}
	return ""
}
