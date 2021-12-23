package main

import (
	"fmt"
	"math"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	solve()
}

const (
	// xMin = 20
	// xMax = 30
	// yMin = -10
	// yMax = -5
	xMin = 281
	xMax = 311
	yMin = -74
	yMax = -54
)

/**
This is just a math problem.

let x = x position
y = y position
a_0 = initial x speed
b_0 = initial y speed

we are trying to choose a_0 and b_0, or find values that work

note that:
x_n = x_(n-1) + a_n
y_n = y_(n-1) + b_n

a_n = a_(n-1) - 1 or a_(n-1) + 1 till a_n = 0
b_n = b_n - 1

if we let n be the value after step n:
x_n = sum(m=0, n-1, a_m) = K_(n-1)
y_n = sum(m=0, n-1, b_m) = L_(n-1)

L_n = sum(n=0, n) b_n = n*b_0 - n(n-1)/2
K_n = ~ looks like a_0 * (a_0 + 1) / 2 if a_0 < n else its a bit less

**/

func solve() {
	// slowest we could go is a speed that after n steps ends just at the target, so we basically solve x_min = (a_0)(a_0+1)/2
	a_min := int(math.Floor(0.5*math.Sqrt(1+4*2*xMin) - 0.5))
	// fastest x speed we could use is just shoot at the target
	a_max := xMax + 1

	// smallest that you could possibly use is just yMin and shoot directly to the target
	b_min := yMin
	// pick an arbitrarily large b_max that we try up to
	b_max := int(math.Abs(yMax)) + 20000

	maxYFound := -1

	velocitiesThatWork := map[string]bool{}

	for a := a_min; a < a_max; a++ {
		for b := b_min; b < b_max; b++ {
			hits, _, maxY := doesSpeedHitTarget(a, b)
			if hits {
				// fmt.Printf("(%d, %d) hits the target after %d steps, max Y=%d\n", a, b, steps, maxY)
				maxYFound = goutil.Max(maxY, maxYFound)
				velocitiesThatWork[fmt.Sprintf("%d,%d", a, b)] = true
			}
		}
	}

	fmt.Println(maxYFound)
	fmt.Println(len(velocitiesThatWork))
}

func doesSpeedHitTarget(a_0, b_0 int) (bool, int, int) {
	n := 0
	x := 0
	y := 0
	a := a_0
	b := b_0

	maxY := 0
	for !(x > xMax || y < yMin) {
		n += 1
		x = x + a
		y = y + b

		if a > 0 {
			a = a - 1
		} else if a < 0 {
			a = a + 1
		} else {
			a = 0
		}

		b = b - 1

		maxY = goutil.Max(maxY, y)

		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			// it hit the spot
			return true, n, maxY
		}
	}
	return false, n, -1
}
