package main

import (
	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.SolutionMain(goutil.NewSolution(func(inputText []string) (int, error) {
		return sol(inputText, 2)
	}, func(inputText []string) (int, error) {
		return sol(inputText, 12)
	}))
}

func sol(inputText []string, k int) (int, error) {
	totalJoltage := 0
	for _, bankString := range inputText {
		bank := newBank(bankString)
		totalJoltage += bank.maxJoltage(k)
	}
	return totalJoltage, nil
}

/*
Initially we tried do do something like trying all possible combinations of prefixes.
That worked in part 1, but failed for part 2 due to the large number of combinations to check.

Instead we realize that we can be rather greedy, as we always want to use the largest number available,
so long as that number doesn't brick us from the rest of the problem.
i.e. if the number ended in 9, we can't start with that 9 as we wouldn't be able to create a 2 or 12 digit number
after that.

We do something greedy then:
- take the highest number available that leaves atleast k-n numbers left in the sequence,
where k is the total number of batteries that need to be selected and n is the current number of selected batteries.
- once that value has been taken, we can reduce our search space to only include values after that index

*/

type Bank struct {
	nums      []int
	batteries map[int]*battery
}

func newBank(bank string) *Bank {
	nums := make([]int, len(bank))
	for i, c := range bank {
		nums[i] = goutil.Atoi(c)
	}

	batteries := map[int]*battery{}

	for i, n := range nums {
		if batteries[n] == nil {
			batteries[n] = newBattery()
		}
		batteries[n].addLocation(i)
	}

	return &Bank{nums, batteries}
}

func (b *Bank) maxJoltage(k int) int {
	joltage := 0
	currIndx := -1

	nextBestJolt := func(x int) (int, int) {
		for jolt := 9; jolt > 0; jolt-- {
			bats := b.batteries[jolt]
			if bats == nil {
				continue
			}
			validLocation := bats.validUse(currIndx, x, len(b.nums))
			if validLocation == -1 {
				// no valid way to use this number, so go down to the next
				// highest joltage
				continue
			}
			// we can use this number, and its always going to be best to use it
			return jolt, validLocation
		}
		return -1, -1
	}

	for x := k - 1; x >= 0; x-- {
		jolt, loc := nextBestJolt(x)
		if jolt == -1 {
			return -1
		}
		joltage += jolt * goutil.PowInt(10, x)
		currIndx = loc
	}

	return joltage
}

type battery struct {
	// locs is already ordered
	locs []int
}

func (b *battery) addLocation(i int) {
	b.locs = append(b.locs, i)
}

func (b *battery) validUse(after, remaining, l int) int {
	for _, loc := range b.locs {
		if loc > after && loc < (l-remaining) {
			return loc
		}
	}
	return -1
}

func newBattery() *battery {
	return &battery{locs: []int{}}
}
