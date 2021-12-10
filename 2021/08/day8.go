package main

import (
	"fmt"
	"sort"
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

func solve1(input []string) error {
	digits := 0
	for _, l := range input {
		parts := strings.Split(l, " | ")

		// signalPatterns := parts[0]
		outputs := parts[1]

		digits += processOutput(outputs)
	}

	fmt.Printf("Solution 1: %d\n", digits)

	return nil
}

func processOutput(output string) int {
	outputDigits := strings.Split(output, " ")

	oFS8s := 0
	for _, o := range outputDigits {

		o = strings.TrimSpace(o)
		outputSignalUnits := len(o)

		switch outputSignalUnits {
		case 2:
			// 1
			oFS8s += 1
		case 4:
			// 4
			oFS8s += 1
		case 3:
			// 7
			oFS8s += 1
		case 7:
			// 8
			oFS8s += 1
		}
	}

	return oFS8s
}

func solve2(input []string) error {
	o := 0
	for _, l := range input {
		parts := strings.Split(l, " | ")

		patterns := strings.Split(strings.TrimSpace(parts[0]), " ")
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")

		output := solveSignal(patterns, outputs)

		o += output
	}

	fmt.Printf("Solution 2: %d\n", o)

	return nil
}

type sig struct {
	top string
	mid string
	bot string

	tl string
	tr string
	bl string
	br string
}

func (s sig) num(p *set) int {
	zero := newSet(s.top + s.tl + s.bl + s.bot + s.br + s.tr)
	one := newSet(s.tr + s.br)
	two := newSet(s.top + s.tr + s.mid + s.bl + s.bot)
	three := newSet(s.top + s.mid + s.bot + s.tr + s.br)
	four := newSet(s.tl + s.mid + s.tr + s.br)
	five := newSet(s.top + s.mid + s.bot + s.tl + s.br)
	six := newSet(s.top + s.tl + s.mid + s.bl + s.br + s.bot)
	seven := newSet(s.top + s.tr + s.br)
	eight := newSet(s.top + s.mid + s.bot + s.bl + s.br + s.tr + s.tl)
	nine := newSet(s.top + s.tr + s.tl + s.mid + s.br + s.bot)

	if p.equal(zero) {
		return 0
	} else if p.equal(one) {
		return 1
	} else if p.equal(two) {
		return 2
	} else if p.equal(three) {
		return 3
	} else if p.equal(four) {
		return 4
	} else if p.equal(five) {
		return 5
	} else if p.equal(six) {
		return 6
	} else if p.equal(seven) {
		return 7
	} else if p.equal(eight) {
		return 8
	} else if p.equal(nine) {
		return 9
	}

	return 1000000
}

func solveSignal(patterns []string, outputs []string) int {
	s := sig{}

	var one *set
	var seven *set
	var four *set
	var eight *set
	for _, p := range patterns {
		if len(p) == 2 {
			one = newSet(p)
		} else if len(p) == 3 {
			seven = newSet(p)
		} else if len(p) == 4 {
			four = newSet(p)
		} else if len(p) == 7 {
			eight = newSet(p)
		}
	}

	// we can find top by doing diff(7, 1)
	s.top = seven.diff(one).String()

	for _, p := range patterns {
		if len(p) == 6 {
			// try some stuff
			if a := four.diff(one).diff(eight.diff(newSet(p))); len(a.s) == 1 {
				// this likely 0
				s.tl = a.String()
			} else if b := one.diff(eight.diff(newSet(p))); len(b.s) == 1 {
				// this likely 6
				s.br = b.String()
			} else {
				// this is likey 9
				s.bl = eight.diff(newSet(p)).String()
			}
		}
	}

	s.tr = one.diff(newSet(s.br)).String()

	s.mid = four.diff(one).diff(newSet(s.tl)).String()

	s.bot = eight.diff(newSet(s.tl + s.top + s.tr + s.br + s.bl + s.mid)).String()

	num := 0
	place := 1000

	for _, o := range outputs {
		num += place * s.num(newSet(o))
		place = place / 10
	}

	fmt.Printf("%+v %v %d\n", s, outputs, num)

	return num
}

type set struct {
	s map[rune]bool
}

func (s *set) String() string {
	o := ""
	for k := range s.s {
		o += string(k)
	}
	return o
}

func newSet(s string) *set {
	st := &set{s: map[rune]bool{}}
	for _, c := range s {
		st.s[c] = true
	}
	return st
}

func (s *set) diff(o *set) *set {
	i, j := s, o

	intersect := map[rune]bool{}

	for k := range i.s {
		intersect[k] = true
	}

	for k := range j.s {
		delete(intersect, k)
	}

	return &set{s: intersect}
}

func (s *set) equal(o *set) bool {
	s1 := SortString(s.String())
	s2 := SortString(o.String())

	return s1 == s2
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
