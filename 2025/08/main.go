package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() { goutil.SolutionMain(goutil.NewSolution(part1, part2)) }

type Box struct {
	X, Y, Z int
	c       *Circuit
}

func NewBox(x, y, z int) *Box {
	return &Box{x, y, z, nil}
}

func (b *Box) dist(o *Box) float64 {
	return goutil.Dist([]int{b.X, b.Y, b.Z}, []int{o.X, o.Y, o.Z})
}

type Circuit struct {
	b []*Box
}

type BoxPair struct {
	b1, b2 *Box
	dist   float64
}

func part1(inputText []string) (int, error) {
	boxes := map[*Box]struct{}{}
	circuits := map[*Circuit]struct{}{}
	for _, l := range inputText {
		parts := strings.Split(l, ",")
		boxes[NewBox(goutil.Atoi(parts[0]), goutil.Atoi(parts[1]), goutil.Atoi(parts[2]))] = struct{}{}
	}

	nClosest := func(n int) []*BoxPair {
		b := make([]*Box, len(boxes))
		i := 0
		for k := range boxes {
			b[i] = k
			i += 1
		}

		pairs := []*BoxPair{}

		for i := 0; i < len(b)-1; i++ {
			for j := i + 1; j < len(b); j++ {
				pairs = append(pairs, &BoxPair{b[i], b[j], float64(b[i].dist(b[j]))})
			}
		}
		slices.SortFunc(pairs, func(a, b *BoxPair) int {
			return cmp.Compare(a.dist, b.dist)
		})
		return pairs[:n]
	}

	mergeCircuits := func(b1, b2 *Box) {
		c := b1.c
		c.b = append(c.b, b2.c.b...)
		cToDel := b2.c
		for _, b := range b2.c.b {
			b.c = c
		}
		delete(circuits, cToDel)
	}

	pairs := nClosest(1000)
	for _, pair := range pairs {
		if pair.b1.c == nil && pair.b2.c == nil {
			// no circuit, so make one
			c := &Circuit{[]*Box{pair.b1, pair.b2}}
			circuits[c] = struct{}{}
			pair.b1.c = c
			pair.b2.c = c
		} else if pair.b1.c != nil && pair.b2.c == nil {
			pair.b1.c.b = append(pair.b1.c.b, pair.b2)
			pair.b2.c = pair.b1.c
		} else if pair.b2.c != nil && pair.b1.c == nil {
			pair.b2.c.b = append(pair.b2.c.b, pair.b1)
			pair.b1.c = pair.b2.c
		} else if pair.b1.c == pair.b2.c {
			// don't need to do anything
			continue
		} else {
			// merge circuit
			mergeCircuits(pair.b1, pair.b2)
		}
	}

	circuitCounts := []int{}
	for c := range circuits {
		circuitCounts = append(circuitCounts, len(c.b))
	}
	slices.Sort(circuitCounts)
	slices.Reverse(circuitCounts)

	return circuitCounts[0] * circuitCounts[1] * circuitCounts[2], nil
}

func part2(inputText []string) (int, error) {
	boxes := map[*Box]struct{}{}
	circuits := map[*Circuit]struct{}{}
	for _, l := range inputText {
		parts := strings.Split(l, ",")
		b := NewBox(goutil.Atoi(parts[0]), goutil.Atoi(parts[1]), goutil.Atoi(parts[2]))
		c := &Circuit{b: []*Box{b}}
		b.c = c
		boxes[b] = struct{}{}
		circuits[c] = struct{}{}
	}

	allPairs := func() []*BoxPair {
		b := make([]*Box, len(boxes))
		i := 0
		for k := range boxes {
			b[i] = k
			i += 1
		}

		pairs := []*BoxPair{}

		for i := 0; i < len(b)-1; i++ {
			for j := i + 1; j < len(b); j++ {
				pairs = append(pairs, &BoxPair{b[i], b[j], float64(b[i].dist(b[j]))})
			}
		}
		slices.SortFunc(pairs, func(a, b *BoxPair) int {
			return cmp.Compare(a.dist, b.dist)
		})
		return pairs
	}

	mergeCircuits := func(b1, b2 *Box) {
		c := b1.c
		c.b = append(c.b, b2.c.b...)
		cToDel := b2.c
		for _, b := range b2.c.b {
			b.c = c
		}
		delete(circuits, cToDel)
	}

	pairs := allPairs()
	for _, pair := range pairs {
		// merge circuit
		if pair.b1.c == pair.b2.c {
			continue
		}
		mergeCircuits(pair.b1, pair.b2)
		if len(circuits) == 1 {
			return pair.b1.X * pair.b2.X, nil
		}
	}
	return -1, nil
}
