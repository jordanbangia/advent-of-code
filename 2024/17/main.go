package main

import (
	"fmt"
	"regexp"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	// input, _ := goutil.ReadFile("test.txt")
	// fmt.Println("Test solutions:")
	// fmt.Println(p1(input))
	// fmt.Println(p2(input))

	input, _ := goutil.ReadFile("input.txt")
	fmt.Println("Input solutions:")
	// fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(inputText []string) string {
	r := regexp.MustCompile(`\d+`)

	// a := goutil.Atoi(r.FindString(inputText[0]))
	// b := goutil.Atoi(r.FindString(inputText[1]))
	// c := goutil.Atoi(r.FindString(inputText[2]))

	r = regexp.MustCompile(`\d`)
	instructionsStr := r.FindAllString(inputText[4], -1)
	instructions := make([]int, len(instructionsStr))
	for i, instr := range instructionsStr {
		instructions[i] = goutil.Atoi(instr)
	}

	return "" //runProgram(a, b, c, instructions)
}

type state struct {
	segs []int64
}

func p2(inputText []string) int64 {
	r := regexp.MustCompile(`\d+`)

	b := goutil.Atoi(r.FindString(inputText[1]))
	c := goutil.Atoi(r.FindString(inputText[2]))

	r = regexp.MustCompile(`\d`)
	instructionsStr := r.FindAllString(inputText[4], -1)
	instructions := make([]int, len(instructionsStr))
	for i, instr := range instructionsStr {
		instructions[i] = goutil.Atoi(instr)
	}

	queue := []state{}
	for i := 0; i < 8; i++ {
		queue = append(queue, state{[]int64{int64(i)}})
	}

	var final int64
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		var x int64
		for i := len(cur.segs) - 1; i >= 0; i-- {
			s := cur.segs[i] << (3 * i)
			x = x | s
		}

		result := runProgram(int(x), b, c, instructions)
		vp := 0
		matched := true
		for p := len(instructions) - len(result); p < len(instructions); p++ {
			if result[vp] != instructions[p] {
				matched = false
				break
			}
			vp++
		}

		done := matched && len(instructions) == len(result)
		if done {
			final = x
			break
		}

		if matched {
			for i := 0; i < 8; i++ {
				nseg := make([]int64, len(cur.segs))
				copy(nseg, cur.segs)
				nseg = append([]int64{int64(i)}, nseg...)
				queue = append(queue, state{nseg})
			}
		}
	}

	return final
}

func runProgram(a, b, c int, instructions []int) []int {
	out := []int{}

	comboOperand := func(op int) int {
		switch op {
		case 0:
			fallthrough
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			return op
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		}
		return -1
	}

	for ip := 0; ip < len(instructions); {
		inst := instructions[ip]
		operand := instructions[ip+1]

		skipIncr := false

		combo := comboOperand(operand)
		switch inst {
		case 0: //adv
			a = a / goutil.PowInt(2, combo)
		case 1: //bxl
			// bitwise xor
			b = b ^ operand
		case 2: //bst
			b = goutil.PosMod(combo, 8)
		case 3: //jnz
			if a != 0 {
				ip = operand
				skipIncr = true
			}
		case 4: //bxc
			b = b ^ c
		case 5: // out
			v := goutil.PosMod(combo, 8)
			out = append(out, v)
		case 6: //bdv
			b = a / goutil.PowInt(2, combo)
		case 7: //cdv
			c = a / goutil.PowInt(2, combo)
		}

		if !skipIncr {
			ip += 2
		}
	}

	return out
}
