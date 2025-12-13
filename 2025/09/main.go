package main

import (
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() { goutil.SolutionMain(goutil.NewSolution(part1, part2)) }

func part1(inputText []string) (int, error) {
	coords := parseCoords(inputText)

	maxArea := 0
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			a := coords[i]
			b := coords[j]
			maxArea = max(maxArea, area(a, b))
		}
	}
	return maxArea, nil
}

func parseCoords(inputText []string) []*goutil.Coord {
	coords := make([]*goutil.Coord, len(inputText))
	for i, l := range inputText {
		parts := strings.Split(l, ",")
		coords[i] = &goutil.Coord{X: goutil.Atoi(parts[0]), Y: goutil.Atoi(parts[1])}
	}
	return coords
}

func area(a, b *goutil.Coord) int {
	return (goutil.Abs(a.X-b.X) + 1) * (goutil.Abs(a.Y-b.Y) + 1)
}

func part2(inputText []string) (int, error) {
	coords := parseCoords(inputText)

	verticalEdges, horizontalEdges := getEdges(coords)

	maxArea := 0
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			a, b := coords[i], coords[j]

			rectArea := area(a, b)
			if rectArea <= maxArea {
				// no point checking area of a rectangle thats smaller than our max
				continue
			}

			// to verify that the rectangle is within the polygon, we test that
			// a point is within the polygon via raycasting
			// and check that there's no edge that intersects the edges of the polygon
			if pointInPolygon(float32(min(a.X, b.X))+0.5, float32(min(a.Y, b.Y))+0.5, verticalEdges) {
				if !polyIntersectsRect(
					min(a.X, b.X),
					min(a.Y, b.Y),
					max(a.X, b.X),
					max(a.Y, b.Y),
					verticalEdges,
					horizontalEdges,
				) {
					maxArea = rectArea
				}
			}
		}
	}

	return maxArea, nil
}

// cast a ray from the point and count how many edges it goes through
func pointInPolygon(pX, pY float32, verticalEdges [][]int) bool {
	intersections := 0

	for _, v := range verticalEdges {
		vx, vymin, vymax := float32(v[0]), float32(v[1]), float32(v[2])

		if vx > pX {
			if vymin <= pY && pY < vymax {
				intersections += 1
			}
		}
	}
	return intersections%2 == 1
}

func getEdges(coords []*goutil.Coord) ([][]int, [][]int) {
	verts := [][]int{}
	horz := [][]int{}

	for i := range coords {
		a := coords[i]
		b := coords[(i+1)%len(coords)]

		if a.X == b.X {
			verts = append(verts, []int{a.X, min(a.Y, b.Y), max(a.Y, b.Y)})
		} else {
			horz = append(horz, []int{min(a.X, b.X), max(a.X, b.X), a.Y})
		}
	}
	return verts, horz
}

func polyIntersectsRect(
	rXmin, rYmin, rXmax, rYmax int,
	verts [][]int,
	horz [][]int,
) bool {
	for _, v := range verts {
		vx, vYmin, vYmax := v[0], v[1], v[2]

		if rXmin < vx && vx < rXmax {
			// edge is stricly within the X range of the rectangle, boundaries are ok
			// does the y part of the rectangle overlap with the y part of the edge
			// if so then this edge intersects into the rectangle
			if max(vYmin, rYmin) < min(rYmax, vYmax) {
				return true
			}
		}
	}

	for _, h := range horz {
		hXmin, hXmax, hY := h[0], h[1], h[2]

		if rYmin < hY && hY < rYmax {
			// similar idea to above
			// does the horizontal edge's Y value overlap the interior of the rectangle
			// if they are at the same height, then check if the the x ranges overlap
			// if they do then this edge is inside of the rectangle
			if max(hXmin, rXmin) < min(hXmax, rXmax) {
				return true
			}
		}
	}
	return false
}
