package main

import (
	"container/heap"
	"errors"
	"fmt"
	"strconv"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	solve1(input)
	solve2(input)
}

func solve1(input []string) {
	grid := make([][]*PathNode, len(input))

	for j, l := range input {
		grid[j] = make([]*PathNode, len(l))
		for i, c := range l {
			v, _ := strconv.Atoi(string(c))
			grid[j][i] = &PathNode{Weight: v, Point: Point{i, j}}
		}
	}

	for y, row := range grid {
		for x, n := range row {
			if y != 0 {
				n.Neighbours = append(n.Neighbours, grid[y-1][x])
			}
			if x != 0 {
				n.Neighbours = append(n.Neighbours, grid[y][x-1])
			}
			if x < len(row)-1 {
				n.Neighbours = append(n.Neighbours, grid[y][x+1])
			}
			if y < len(grid)-1 {
				n.Neighbours = append(n.Neighbours, grid[y+1][x])
			}
		}
	}

	// at this point, the grid has been constructed

	start := grid[0][0]
	maxX, maxY := len(grid[0])-1, len(grid)-1
	end := grid[maxY][maxX]

	path, err := findPath(start, end)
	if err != nil {
		fmt.Println("whoops")
	}
	fmt.Printf("Solution 1: %d\n", risk(path, grid))
}

func solve2(input []string) {
	iGrid := make([][]*PathNode, len(input))

	for j, l := range input {
		iGrid[j] = make([]*PathNode, len(l))
		for i, c := range l {
			v, _ := strconv.Atoi(string(c))
			iGrid[j][i] = &PathNode{Weight: v, Point: Point{i, j}}
		}
	}

	grid := make([][]*PathNode, len(iGrid)*5)
	for i := range grid {
		grid[i] = make([]*PathNode, len(iGrid[0])*5)
	}

	m := len(iGrid)
	n := len(iGrid[0])

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y, row := range iGrid {
				for x, p := range row {
					weight := (p.Weight + i + j) % 9
					if weight == 0 {
						weight = 9
					}
					grid[y+m*j][x+n*i] = &PathNode{
						Weight: weight,
						Point:  Point{x + n*i, y + m*j},
					}
				}
			}
		}
	}

	for y, row := range grid {
		for x, n := range row {
			if y != 0 {
				n.Neighbours = append(n.Neighbours, grid[y-1][x])
			}
			if x != 0 {
				n.Neighbours = append(n.Neighbours, grid[y][x-1])
			}
			if x < len(row)-1 {
				n.Neighbours = append(n.Neighbours, grid[y][x+1])
			}
			if y < len(grid)-1 {
				n.Neighbours = append(n.Neighbours, grid[y+1][x])
			}
		}
	}

	start := grid[0][0]
	maxX, maxY := len(grid[0])-1, len(grid)-1
	end := grid[maxY][maxX]

	path, err := findPath(start, end)
	if err != nil {
		fmt.Println("whoops")
	}
	fmt.Printf("Solution 2: %d\n", risk(path, grid))
}

type PathNode struct {
	Point
	Weight     int
	PathScore  int
	Neighbours []*PathNode
}

type Point struct {
	x, y int
}

func findPath(start, end *PathNode) ([]Point, error) {
	q := PathQueue{}

	heap.Init(&q)
	heap.Push(&q, start)

	cameFrom := map[Point]Point{}
	gScore := map[Point]int{start.Point: 0}
	for q.Len() > 0 {
		cur := heap.Pop(&q).(*PathNode)
		if cur.Point == end.Point {
			return makePath(end, cameFrom), nil
		}
		for _, n := range cur.Neighbours {
			gs := gScore[cur.Point] + n.Weight
			if gScore[n.Point] == 0 || gs < gScore[n.Point] {
				cameFrom[n.Point] = cur.Point
				gScore[n.Point] = gs
				newNode := &PathNode{
					Point:      n.Point,
					Weight:     n.Weight,
					PathScore:  gs + (end.x - n.x) + (end.y - n.y),
					Neighbours: n.Neighbours,
				}
				heap.Push(&q, newNode)
			}
		}
	}

	return nil, errors.New("whoops")
}

func makePath(end *PathNode, cameFrom map[Point]Point) []Point {
	p := []Point{end.Point}
	nextNode := cameFrom[end.Point]
	start := Point{0, 0}
	for nextNode != start {
		p = append(p, nextNode)
		nextNode = cameFrom[nextNode]
	}
	return p
}

func risk(p []Point, g [][]*PathNode) int {
	t := 0
	for _, o := range p {
		t += g[o.y][o.x].Weight
	}
	return t
}

type PathQueue []*PathNode

func (pq PathQueue) Len() int {
	return len(pq)
}

func (pq PathQueue) Less(i, j int) bool {
	return pq[i].PathScore < pq[j].PathScore
}

func (pq PathQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PathQueue) Push(x interface{}) {
	n := x.(*PathNode)
	*pq = append(*pq, n)
}

func (pq *PathQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return node
}
