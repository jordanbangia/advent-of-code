package main

import (
	"fmt"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewOneFuncSolution(solStupid),
		false,
	)
}

const dialStart = 50

func sol(text []string) (int, int, error) {
	place := dialStart
	endZeroes := 0
	crossZeroes := 0
	for _, line := range text {
		dir := line[0]
		num := goutil.Atoi(line[1:])
		fullRotations := num / 100
		end := place

		if dir == 'L' {
			end -= num

			toZero := place
			if num >= toZero && place != 0 {
				num = num - toZero
				crossZeroes += 1 + fullRotations
				// fmt.Println("cross zeroes: ", crossZeroes)
			}
		} else {
			end += num

			toHundred := 100 - place
			if num >= toHundred && place != 0 {
				num = num - toHundred
				crossZeroes += 1 + fullRotations
				// fmt.Println("cross zeroes: ", crossZeroes)
			}
		}

		fmt.Println(place, line, (100+end)%100, "-", crossZeroes)
		place = (100 + end) % 100
		if place == 0 {
			endZeroes += 1
		}
		// fmt.Println("cross zero:", crossZeroes, "end zero:", endZeroes)

	}
	fmt.Println(place)
	return endZeroes, crossZeroes, nil
}

func solStupid(text []string) (int, int, error) {
	place := dialStart

	endZeroes := 0
	crossZeroes := 0

	for _, line := range text {
		dir := line[0]
		num := goutil.Atoi(line[1:])

		step := 1
		if dir == 'L' {
			step = -1
		}

		for range num {
			place += step
			if place < 0 {
				place += 100
			}
			place = place % 100
			if place == 0 {
				crossZeroes += 1
			}
		}
		if place == 0 {
			endZeroes += 1
		}
	}
	return endZeroes, crossZeroes, nil
}
