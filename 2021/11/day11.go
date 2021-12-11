package main

import (
	"fmt"
	"strconv"
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

func parseMap(input []string) [][]int {
	hMap := make([][]int, len(input))

	for i, l := range input {
		row := make([]int, len(l))
		for j, c := range l {
			row[j], _ = strconv.Atoi(string(c))
		}
		hMap[i] = row
	}
	return hMap
}

func solve1(input []string) error {
	m := parseMap(input)

	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += doStep(m)
		fmt.Printf("Done step %d: %d\n", i+1, totalFlashes)
		// fmt.Printf("%+v\n", m)
	}

	fmt.Println(totalFlashes)

	return nil
}

func solve2(input []string) error {
	m := parseMap(input)

	totalFlashes := 0
	i := 0
	for {
		flashesDuringStep := doStep(m)
		totalFlashes += flashesDuringStep
		fmt.Printf("Done step %d: %d\n", i+1, totalFlashes)
		if flashesDuringStep == len(m)*len(m[0]) {
			fmt.Println(i + 1)
			return nil
		}
		i += 1
	}

	return nil
}

func doStep(mp [][]int) int {
	m := len(mp)
	n := len(mp[0])

	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[i][j] += 1
			if mp[i][j] > 9 {
				q = append(q, []int{i, j})
			}
		}
	}

	// fmt.Println(mp)
	// fmt.Println(q)

	hasFlashed := map[string]bool{}

	for len(q) > 0 {
		var a []int
		a, q = q[0], q[1:]

		x, y := a[0], a[1]
		if mp[x][y] > 9 && !hasFlashed[key(x, y)] {
			hasFlashed[key(x, y)] = true
			q = append(q, addOneAround(mp, x, y)...)
		}
	}

	for k := range hasFlashed {
		x, y := parseKey(k)
		mp[x][y] = 0
	}

	return len(hasFlashed)
}

func addOneAround(mp [][]int, x, y int) [][]int {
	updates := [][]int{}
	if x > 0 {
		updates = append(updates, []int{x - 1, y})
	}
	if x < len(mp)-1 {
		updates = append(updates, []int{x + 1, y})
	}
	if y > 0 {
		updates = append(updates, []int{x, y - 1})
	}
	if y < len(mp[0])-1 {
		updates = append(updates, []int{x, y + 1})
	}
	if x > 0 && y > 0 {
		updates = append(updates, []int{x - 1, y - 1})
	}
	if x < len(mp)-1 && y > 0 {
		updates = append(updates, []int{x + 1, y - 1})
	}
	if x > 0 && y < len(mp[0])-1 {
		updates = append(updates, []int{x - 1, y + 1})
	}
	if x < len(mp)-1 && y < len(mp[0])-1 {
		updates = append(updates, []int{x + 1, y + 1})
	}

	ret := [][]int{}
	for _, u := range updates {
		mp[u[0]][u[1]] += 1
		if mp[u[0]][u[1]] > 9 {
			ret = append(ret, u)
		}
	}

	return ret
}

func key(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func parseKey(s string) (int, int) {
	p := strings.Split(s, "_")
	xS := p[0]
	yS := p[1]

	x, _ := strconv.Atoi(xS)
	y, _ := strconv.Atoi(yS)
	return x, y
}
