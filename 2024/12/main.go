package main

import (
	"fmt"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			nil, p2,
		),
		true,
	)
}

type Region struct {
	label     string
	plots     map[string]struct{}
	edgePlots map[string]map[string]struct{}
	edges     int
	sides     int
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func findPlots(g map[string]string) []Region {
	possibleRegions := goutil.NewStack[[]int]()
	possibleRegions.Push([]int{0, 0})

	visited := map[string]struct{}{}

	regions := []Region{}

	for !possibleRegions.IsEmpty() {
		start, _ := possibleRegions.Pop()

		// don't check nodes that don't exist
		if _, exists := g[goutil.AKey(start)]; !exists {
			continue
		}

		// don't check nodes we've already checked
		if _, hasVisited := visited[goutil.AKey(start)]; hasVisited {
			continue
		}

		// do a breadth first search from our starting node
		region := Region{label: g[goutil.AKey(start)]}
		visitedInRegion := map[string]struct{}{}
		edgePlots := map[string]map[string]struct{}{}
		s := goutil.NewStack[[]int]()
		s.Push(start)

		for !s.IsEmpty() {
			x, _ := s.Pop()
			if _, xIsVisited := visitedInRegion[goutil.AKey(x)]; xIsVisited {
				continue
			}
			visitedInRegion[goutil.AKey(x)] = struct{}{}

			edges := 0
			for _, d := range dirs {
				n := []int{x[0] + d[0], x[1] + d[1]}
				if _, alreadyVisited := visitedInRegion[goutil.AKey(n)]; alreadyVisited {
					continue
				}
				l, exists := g[goutil.AKey(n)]
				if !exists {
					if _, exists := edgePlots[goutil.AKey(x)]; !exists {
						edgePlots[goutil.AKey(x)] = map[string]struct{}{}
					}

					edgePlots[goutil.AKey(x)][goutil.AKey(d)] = struct{}{}
					edges += 1
					continue
				}
				if l != region.label {
					if _, exists := edgePlots[goutil.AKey(x)]; !exists {
						edgePlots[goutil.AKey(x)] = map[string]struct{}{}
					}
					edgePlots[goutil.AKey(x)][goutil.AKey(d)] = struct{}{}
					edges += 1
					possibleRegions.Push(n)
				} else {
					s.Push(n)
				}
			}
			region.edges += edges
		}

		for k := range visitedInRegion {
			visited[k] = struct{}{}
		}

		region.plots = visitedInRegion
		region.edgePlots = edgePlots
		regions = append(regions, region)
	}
	return regions
}

func p1(inputText []string) (int, error) {
	g := map[string]string{}

	for i, l := range inputText {
		for j, c := range l {
			g[goutil.Key(i, j)] = string(c)
		}
	}

	cost := 0
	for _, r := range findPlots(g) {
		cost += r.edges * len(r.plots)
	}

	return cost, nil
}

func p2(inputText []string) (int, error) {
	g := map[string]string{}

	for i, l := range inputText {
		for j, c := range l {
			g[goutil.Key(i, j)] = string(c)
		}
	}

	plots := findPlots(g)

	price := 0
	for _, plot := range plots {
		sides := 0

		for e, edgeDirs := range plot.edgePlots {
			if len(edgeDirs) == 0 {
				continue
			}

			for d := range edgeDirs {
				// d says that from e to d there is an edge
				// go 90deg in both directions, and try to match other edges
				dVal0, _ := goutil.SplitKey(d)
				dir1 := []int{}
				dir2 := []int{}
				if dVal0 == 0 {
					dir1 = []int{1, 0}
					dir2 = []int{-1, 0}
				} else {
					dir1 = []int{0, 1}
					dir2 = []int{0, -1}
				}

				for _, travelDir := range [][]int{dir1, dir2} {
					r, c := goutil.SplitKey(e)
					r += travelDir[0]
					c += travelDir[1]
					for {
						edge, exists := plot.edgePlots[goutil.Key(r, c)]
						if !exists {
							// position is not an edge, so we broken the chain
							break
						}
						if _, hasSameEdge := edge[d]; hasSameEdge {
							delete(plot.edgePlots[goutil.Key(r, c)], d)
							r += travelDir[0]
							c += travelDir[1]
						} else {
							break
						}
					}
				}
				delete(plot.edgePlots[e], d)
				sides += 1
			}
		}

		fmt.Printf("%s has %d sides\n", plot.label, sides)
		price += sides * len(plot.plots)
	}
	return price, nil
}
