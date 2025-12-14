package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() { goutil.SolutionMain(goutil.NewSolution(part1, part2)) }

func part1(inputText []string) (int, error) {
	solveProblem := func(p *Problem) int {
		type S struct {
			currentLightState int
			lastButtonPress   int // index in p.buttons
			presses           int
		}

		buttonToInt := func(b []int) int {
			d := 0
			for _, i := range b {
				d = d | 1<<i
			}
			return d
		}

		q := goutil.NewQueue[S]()
		q.Push(S{0, -1, 0})

		seenStates := map[int]int{}

		for !q.IsEmpty() {
			n, _ := q.Pop()

			pressesToCurrentState, exists := seenStates[n.currentLightState]
			if exists && pressesToCurrentState < n.presses {
				// we've already been to this light state with
				// fewere button presses, we can drop this and move on
				continue
			}
			seenStates[n.currentLightState] = n.presses

			for i, b := range p.buttons {
				if i == n.lastButtonPress {
					// avoid pressing the same button twice in a row
					continue
				}
				nextState := n.currentLightState ^ buttonToInt(b)
				if nextState == p.desiredState {
					return n.presses + 1
				}
				q.Push(S{nextState, i, n.presses + 1})
			}
		}
		return -1
	}

	answer := 0
	for _, l := range inputText {
		problem := parseLine(l)
		fewestPresses := solveProblem(problem)
		fmt.Println(problem.String(), fewestPresses)
		answer += fewestPresses
	}

	return answer, nil
}

func part2(inputText []string) (int, error) {
	answers := make(chan int, len(inputText))
	wg := &sync.WaitGroup{}
	for _, l := range inputText {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			problem := parseLine(l)
			buttonMap := map[int][]int{}
			for i, b := range problem.buttons {
				buttonMap[i] = b
			}

			a := solve2(problem.joltageRequirements, buttonMap)
			fmt.Println(problem.String(), a)
			answers <- a
		}(l)
	}

	wg.Wait()
	close(answers)

	fewestPresses := 0
	for a := range answers {
		fewestPresses += a
	}

	return fewestPresses, nil
}

type Problem struct {
	desiredState        int
	buttons             [][]int
	joltageRequirements []int
}

func (p *Problem) String() string {
	return fmt.Sprintf(
		"%b | %+v | %+v", p.desiredState, p.buttons, p.joltageRequirements,
	)
}

func parseLine(line string) *Problem {
	desiredStateRegex := regexp.MustCompile(`\[([\.#]*)\]`)
	desiredStates := desiredStateRegex.FindAllStringSubmatch(line, -1)
	if len(desiredStates) != 1 {
		panic(desiredStates)
	}

	buttonRegex := regexp.MustCompile(`\(([\d,?]*)\)`)
	buttons := buttonRegex.FindAllStringSubmatch(line, -1)

	joltageRegex := regexp.MustCompile(`\{([\d,?]*)\}`)
	joltage := joltageRegex.FindAllStringSubmatch(line, -1)
	if len(joltage) != 1 {
		panic(joltage)
	}

	p := &Problem{
		desiredState:        parseDesiredState(desiredStates[0][1]),
		buttons:             stringArrayToArrayOfIntArray(buttons),
		joltageRequirements: stringToIntArray(joltage[0][1]),
	}

	return p
}

func parseDesiredState(s string) int {
	d := 0
	for k := 0; k < len(s); k++ {
		if s[k] == '#' {
			d = d | 1<<k
		}
	}
	return d
}

func stringArrayToArrayOfIntArray(s [][]string) [][]int {
	v := make([][]int, len(s))
	for i, l := range s {
		v[i] = stringToIntArray(l[1])
	}
	return v
}

func stringToIntArray(s string) []int {
	p := strings.Split(s, ",")
	i := make([]int, len(p))
	for j, k := range p {
		i[j] = goutil.Atoi(k)
	}
	return i
}

func allEqual[T comparable](a []T, b T) bool {
	for _, k := range a {
		if k != b {
			return false
		}
	}
	return true
}

func solve2(joltage []int, availableButtons map[int][]int) int {
	if allEqual(joltage, 0) {
		return 0
	}

	joltageIndxWithLeastOptions, buttonsThatEffectJoltage := func() (int, []int) {
		buttonsThatEffectJoltage := map[int][]int{}
		for i, b := range availableButtons {
			for _, j := range b {
				buttonsThatEffectJoltage[j] = append(buttonsThatEffectJoltage[j], i)
			}
		}

		joltageWithFewestOptions := -1
		fewestOptions := -1
		for j, options := range buttonsThatEffectJoltage {
			if joltageWithFewestOptions == -1 {
				joltageWithFewestOptions = j
				fewestOptions = len(options)
			} else {
				if len(options) < fewestOptions {
					joltageWithFewestOptions = j
					fewestOptions = len(options)
				} else if len(options) == fewestOptions {
					if joltage[j] > joltage[joltageWithFewestOptions] {
						// want the one with the highest joltage requirement, it will eliminate the most options
						joltageWithFewestOptions = j
					}
				}
			}
		}
		return joltageWithFewestOptions, buttonsThatEffectJoltage[joltageWithFewestOptions]
	}()

	presses := joltage[joltageIndxWithLeastOptions]

	results := 1_000_000_000_000

	// there are some buttons that we could hit to effect the joltage
	// this is the smallest set of buttons that we could hit
	if len(buttonsThatEffectJoltage) > 0 {
		// make an array of counts, representing the partition of the button presses
		counts := make([]int, len(buttonsThatEffectJoltage))

		// start with pressing one of them the full required amount of times,
		// we'll progressively move the presses around
		counts[len(buttonsThatEffectJoltage)-1] = presses

		for {
			newJoltage := goutil.Duplicate(joltage)
			good := true

		buttons:
			for b, cnt := range counts {
				// for each button that we're pressing
				if cnt == 0 {
					// if we're not pressing it continue
					continue
				}
				// that button potentially affects other joltages
				// so we check that if the press would leave us with too much joltage
				// (or in our case negative since we're subtracting)
				buttonToPress := availableButtons[buttonsThatEffectJoltage[b]]
				for _, j := range buttonToPress {
					if newJoltage[j] >= cnt {
						newJoltage[j] -= cnt
					} else {
						// we can't do this state of state of presses at all
						// as it would leave us in an invalid state
						good = false
						break buttons
					}
				}
			}

			// we can do the intended button press combo
			// so we remove those buttons from the available buttons
			// as any more would cause us to go over for the intended joltage
			// and then we run the algorithm again for the reduced state
			if good {
				nextAvailableButtons := goutil.CopyMap(availableButtons)
				for _, b := range buttonsThatEffectJoltage {
					delete(nextAvailableButtons, b)
				}
				r := solve2(newJoltage, nextAvailableButtons)
				if r != 1_000_000_000_000 {
					results = min(r+presses, results)
				}
			}

			hasMore := nextCombination(counts)
			if !hasMore {
				break
			}
		}
	}
	return results
}

func nextCombination(c []int) bool {
	rightMost := 0
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] > 0 {
			rightMost = i
			break
		}
	}
	if rightMost == 0 {
		return false
	}
	k := c[rightMost]
	c[rightMost-1] += 1
	c[rightMost] = 0
	c[len(c)-1] = k - 1
	return true
}
