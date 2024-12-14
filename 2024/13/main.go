package main

import (
	"fmt"
	"math"
	"regexp"

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

type Game struct {
	aX, aY, bX, bY, tX, tY float64
}

func solveGame(g *Game) int {
	fmt.Printf("%+v\n", g)

	a := (g.tX - g.tY*(g.bX/g.bY)) / (g.aX - g.aY*(g.bX/g.bY))
	fmt.Println(a)
	if math.Abs(a-math.Round(a)) < 1e-4 {
		a = math.Round(a)
	} else {
		fmt.Println("Can't reach the target")
		return -1
	}
	b := (g.tY - g.aY*a) / g.bY

	c := 3*a + b

	fmt.Printf("Can win with %d tokens\n", int(c))
	return int(c)
}

func countTokens(inputText []string, offset int) (int, error) {
	rX := regexp.MustCompile(`X\+(\d*)`)
	rY := regexp.MustCompile(`Y\+(\d*)`)
	tX := regexp.MustCompile(`X=(\d*)`)
	tY := regexp.MustCompile(`Y=(\d*)`)

	tokens := 0
	for i := 0; i < len(inputText); i += 4 {
		g := &Game{}

		buttonALine := inputText[i]
		buttonBLine := inputText[i+1]
		target := inputText[i+2]

		g.aX = float64(goutil.Atoi(rX.FindAllStringSubmatch(buttonALine, -1)[0][1]))
		g.aY = float64(goutil.Atoi(rY.FindAllStringSubmatch(buttonALine, -1)[0][1]))

		g.bX = float64(goutil.Atoi(rX.FindAllStringSubmatch(buttonBLine, -1)[0][1]))
		g.bY = float64(goutil.Atoi(rY.FindAllStringSubmatch(buttonBLine, -1)[0][1]))

		g.tX = float64(goutil.Atoi(tX.FindAllStringSubmatch(target, -1)[0][1])) + float64(offset)
		g.tY = float64(goutil.Atoi(tY.FindAllStringSubmatch(target, -1)[0][1])) + float64(offset)

		if s := solveGame(g); s != -1 {
			tokens += s
		}
	}
	return tokens, nil
}

func p1(inputText []string) (int, error) {
	return countTokens(inputText, 0)
}

func p2(inputText []string) (int, error) {
	return countTokens(inputText, 10000000000000)
}
