package main

import (
	"fmt"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {

	files := []string{
		"test_input_1.txt",
		"test_input_2.txt",
		"test_input_3.txt",
		"input.txt",
	}

	for _, f := range files {
		input, err := goutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("--------")
		fmt.Printf("Solutions for file: %s\n", f)
		_ = solve1(input)
		_ = solve2(input)
	}

}

type Node struct {
	name string
	adj  []*Node

	isSmall bool
}

func NewNode(name string) *Node {
	return &Node{
		name:    name,
		adj:     []*Node{},
		isSmall: strings.ToLower(name) == name,
	}
}

func (n *Node) addAdjacency(o *Node) {
	n.adj = append(n.adj, o)
}

func solve1(input []string) error {
	nodeMap := buildMap(input)

	// do a search to count paths
	paths := countPaths(nodeMap, "start", map[string]struct{}{}, 0)

	fmt.Printf("Solution 1: %d\n", paths)

	return nil
}

func solve2(input []string) error {
	nodeMap := buildMap(input)

	// do a search to count paths
	paths := coutPathsOneReturn(nodeMap, "start", map[string]int{}, "")

	fmt.Printf("Solution 2: %d\n", paths)

	return nil
}

func buildMap(input []string) map[string]*Node {
	nodeMap := map[string]*Node{}

	// create nodes
	for _, l := range input {
		nodes := strings.Split(l, "-")
		for _, n := range nodes {
			nodeMap[n] = NewNode(n)
		}
	}

	// apply adjacencies
	for _, l := range input {
		nodes := strings.Split(l, "-")
		n := nodes[0]
		o := nodes[1]

		nodeMap[n].addAdjacency(nodeMap[o])
		nodeMap[o].addAdjacency(nodeMap[n])
	}
	return nodeMap
}

func countPaths(nm map[string]*Node, next string, visited map[string]struct{}, level int) int {
	n := nm[next]

	visited[next] = struct{}{}

	// fmt.Println(next, visited)

	c := 0

	adjacents := n.adj
	for _, o := range adjacents {
		if _, ok := visited[o.name]; ok && o.isSmall {
			// we've already visited this small node, don't visit it again
			continue
		}

		if o.name == "end" {
			// one of the adjacent nodes is the exit, so just record that we could make it to the end
			c++
		} else {
			c += countPaths(nm, o.name, copyMap(visited), level+1)
		}
	}

	return c
}

func copyMap(v map[string]struct{}) map[string]struct{} {
	n := map[string]struct{}{}

	for k := range v {
		n[k] = struct{}{}
	}
	return n
}

func copyMapI(v map[string]int) map[string]int {
	n := map[string]int{}

	for k, o := range v {
		n[k] = o
	}
	return n
}

func coutPathsOneReturn(nm map[string]*Node, next string, visited map[string]int, path string) int {
	n := nm[next]

	visited[next] += 1

	// fmt.Println(next, visited)

	c := 0

	adjacents := n.adj
	for _, o := range adjacents {
		if o.name == "end" {
			// one of the adjacent nodes is the exit, so just record that we could make it to the end
			c++
			// fmt.Println(path + "," + next + "," + o.name)
		} else if canVisit(visited, nm, o) {
			c += coutPathsOneReturn(nm, o.name, copyMapI(visited), path+","+next)
		}
	}

	return c
}

func canVisit(visited map[string]int, nm map[string]*Node, o *Node) bool {
	if o.name == "start" {
		// can only come from start once
		return false
	}

	if visited[o.name] == 0 {
		return true
	}
	if visited[o.name] > 0 && !o.isSmall {
		return true
	}
	// we've seen this node before and its a small node
	for k, n := range visited {
		if n >= 2 && nm[k].isSmall {
			// some small node has already been visited twice, so we can't visit any more nodes twice
			return false
		}
	}
	return true
}
