package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	goutil.RunSolution(
		goutil.NewSolution(
			p1, p2,
		),
		false,
	)
}

type MapEntry struct {
	id    int
	size  int
	moved bool
}

func (e *MapEntry) String() string {
	return fmt.Sprintf("%d: %d", e.id, e.size)
}

func printDisk(m []MapEntry) string {
	sb := strings.Builder{}
	for _, e := range m {
		c := "."
		if e.id != -1 {
			c = strconv.Itoa(e.id)
		}
		for i := 0; i < e.size; i++ {
			sb.WriteString(c)
		}
	}
	return sb.String()
}

func p1(inputText []string) (int, error) {
	inputLine := inputText[0]

	mapEntries := []MapEntry{}

	ids := 0
	isFile := true
	for i := 0; i < len(inputLine); i += 1 {
		if isFile {
			mapEntries = append(mapEntries, MapEntry{id: ids, size: goutil.Atoi(string(inputLine[i]))})
			ids += 1
		} else {
			mapEntries = append(mapEntries, MapEntry{id: -1, size: goutil.Atoi(string(inputLine[i]))})
		}
		isFile = !isFile
	}

	// fmt.Printf("%s\n", printDisk(mapEntries))

	endElement := func() int {
		for j := len(mapEntries) - 1; j >= 0; j -= 1 {
			if mapEntries[j].id != -1 {
				return j
			}
		}
		return -1
	}

	for i := 0; i < len(mapEntries); i++ {
		// skip over items that are already placed at the front
		if mapEntries[i].id != -1 {
			continue
		}

		// we've found an empty spot, grab an item from the back and try to insert it into here
		j := endElement()
		if j == -1 {
			continue
		}
		if j < i {
			break
		}

		switch {
		case mapEntries[i].size == mapEntries[j].size:
			// exact swap
			mapEntries[i].id = mapEntries[j].id
			mapEntries[j].id = -1
		case mapEntries[i].size > mapEntries[j].size:
			// our empty spot has more space than necessary for the data we need to fill
			mapEntries[i].id = mapEntries[j].id
			mapEntries[j].id = -1
			leftOver := mapEntries[i].size - mapEntries[j].size
			mapEntries[i].size = mapEntries[j].size
			mapEntries = slices.Insert(mapEntries, i+1, MapEntry{id: -1, size: leftOver})
		case mapEntries[i].size < mapEntries[j].size:
			// our empty spot doesn't have enough space for all of the item from the back
			// use up as much as we can
			mapEntries[i].id = mapEntries[j].id
			mapEntries[j].size = mapEntries[j].size - mapEntries[i].size
		}
		// fmt.Printf("%s\n", printDisk(mapEntries))
	}

	return checksum(mapEntries), nil
}

func checksum(mapEntries []MapEntry) int {
	k := 0
	s := 0
	for _, entry := range mapEntries {
		// fmt.Println(entry)
		if entry.id != -1 {
			for x := 0; x < entry.size; x++ {
				s += (k + x) * entry.id
				// fmt.Println(s)
			}
		}
		k += entry.size
	}
	return s
}

func p2(inputText []string) (int, error) {
	inputLine := inputText[0]

	mapEntries := []MapEntry{}

	ids := 0
	isFile := true
	for i := 0; i < len(inputLine); i += 1 {
		if isFile {
			mapEntries = append(mapEntries, MapEntry{id: ids, size: goutil.Atoi(string(inputLine[i]))})
			ids += 1
		} else {
			mapEntries = append(mapEntries, MapEntry{id: -1, size: goutil.Atoi(string(inputLine[i]))})
		}
		isFile = !isFile
	}

	// fmt.Println(printDisk(mapEntries))

	// this one should be simpler, we're just swapping
	for i := len(mapEntries) - 1; i >= 0; i -= 1 {
		// start from the last element
		// skip blocks that are empty or have already moved
		if mapEntries[i].id == -1 || mapEntries[i].moved {
			continue
		}

		// we've found an entry at this point that could be moved
		// try to find a spot that is avialbe that we can move it to
		for j := 0; j < i; j++ {
			if mapEntries[j].id == -1 {
				if mapEntries[j].size == mapEntries[i].size {
					mapEntries[j].id = mapEntries[i].id
					mapEntries[i].id = -1
					mapEntries[j].moved = true
					break
				} else if mapEntries[j].size > mapEntries[i].size {
					// more space than necssary
					mapEntries[j].id, mapEntries[i].id = mapEntries[i].id, -1
					mapEntries[j].moved = true
					leftOver := mapEntries[j].size - mapEntries[i].size
					mapEntries[j].size = mapEntries[i].size
					mapEntries = slices.Insert(mapEntries, j+1, MapEntry{id: -1, size: leftOver})
					break
				}
			}
		}
		// fmt.Println(printDisk(mapEntries))
	}

	return checksum(mapEntries), nil
}
