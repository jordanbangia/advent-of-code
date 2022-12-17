package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	// part2(SampleSensor, 0, 20)
	part2(ProdSensors, 0, 4000000)
}

func part2(sensors []*Sensor, minCoord, maxCoord int) {
	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan Sensor)

	work := make(chan int, 4000000)

	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range work {
				select {
				case <-ctx.Done():
					return
				default:
				}
				fmt.Println("Starting", y)
				for x := minCoord; x < maxCoord; x++ {
					if isOutsideAllSensors(x, y, sensors) {
						fmt.Println(fmt.Sprintf("%d, %d", x, y))
						result <- Sensor{X: x, Y: y}
						return
					}
				}
				fmt.Println("Finishing", y)
			}
		}()
	}

	fmt.Println("creating work")
	for i := minCoord; i < maxCoord; i++ {
		work <- i
	}
	close(work)
	fmt.Println("created work")

	r := <-result
	cancel()
	wg.Wait()
	fmt.Println(r)
}

func isOutsideAllSensors(x, y int, sensors []*Sensor) bool {
	for _, sensor := range sensors {
		if dist(x, y, sensor.X, sensor.Y) <= sensor.dist() {
			return false
		}
	}
	return true
}

type Beacon struct {
	X int
	Y int
}

type Sensor struct {
	X int
	Y int

	B *Beacon
}

func (s *Sensor) dist() int {
	return goutil.Abs(s.X-s.B.X) + goutil.Abs(s.Y-s.B.Y)
}

func dist(x1, y1, x2, y2 int) int {
	return goutil.Abs(x1-x2) + goutil.Abs(y1-y2)
}

var SampleSensor = []*Sensor{
	{X: 2, Y: 18, B: &Beacon{X: -2, Y: 15}},
	{X: 9, Y: 16, B: &Beacon{X: 10, Y: 16}},
	{X: 13, Y: 2, B: &Beacon{X: 15, Y: 3}},
	{X: 12, Y: 14, B: &Beacon{X: 10, Y: 16}},
	{X: 10, Y: 20, B: &Beacon{X: 10, Y: 16}},
	{X: 14, Y: 17, B: &Beacon{X: 10, Y: 16}},
	{X: 8, Y: 7, B: &Beacon{X: 2, Y: 10}},
	{X: 2, Y: 0, B: &Beacon{X: 2, Y: 10}},
	{X: 0, Y: 11, B: &Beacon{X: 2, Y: 10}},
	{X: 20, Y: 14, B: &Beacon{X: 25, Y: 17}},
	{X: 17, Y: 20, B: &Beacon{X: 21, Y: 22}},
	{X: 16, Y: 7, B: &Beacon{X: 15, Y: 3}},
	{X: 14, Y: 3, B: &Beacon{X: 15, Y: 3}},
	{X: 20, Y: 1, B: &Beacon{X: 15, Y: 3}},
}

var ProdSensors = []*Sensor{
	{X: 3729579, Y: 1453415, B: &Beacon{X: 4078883, Y: 2522671}},
	{X: 3662668, Y: 2749205, B: &Beacon{X: 4078883, Y: 2522671}},
	{X: 257356, Y: 175834, B: &Beacon{X: 1207332, Y: 429175}},
	{X: 2502777, Y: 3970934, B: &Beacon{X: 3102959, Y: 3443573}},
	{X: 24076, Y: 2510696, B: &Beacon{X: 274522, Y: 2000000}},
	{X: 3163363, Y: 3448163, B: &Beacon{X: 3102959, Y: 3443573}},
	{X: 1011369, Y: 447686, B: &Beacon{X: 1207332, Y: 429175}},
	{X: 3954188, Y: 3117617, B: &Beacon{X: 4078883, Y: 2522671}},
	{X: 3480746, Y: 3150039, B: &Beacon{X: 3301559, Y: 3383795}},
	{X: 2999116, Y: 3137910, B: &Beacon{X: 3102959, Y: 3443573}},
	{X: 3546198, Y: 462510, B: &Beacon{X: 3283798, Y: -405749}},
	{X: 650838, Y: 1255586, B: &Beacon{X: 274522, Y: 2000000}},
	{X: 3231242, Y: 3342921, B: &Beacon{X: 3301559, Y: 3383795}},
	{X: 1337998, Y: 31701, B: &Beacon{X: 1207332, Y: 429175}},
	{X: 1184009, Y: 3259703, B: &Beacon{X: 2677313, Y: 2951659}},
	{X: 212559, Y: 1737114, B: &Beacon{X: 274522, Y: 2000000}},
	{X: 161020, Y: 2251470, B: &Beacon{X: 274522, Y: 2000000}},
	{X: 3744187, Y: 3722432, B: &Beacon{X: 3301559, Y: 3383795}},
	{X: 2318112, Y: 2254019, B: &Beacon{X: 2677313, Y: 2951659}},
	{X: 2554810, Y: 56579, B: &Beacon{X: 3283798, Y: -405749}},
	{X: 1240184, Y: 897870, B: &Beacon{X: 1207332, Y: 429175}},
	{X: 2971747, Y: 2662873, B: &Beacon{X: 2677313, Y: 2951659}},
	{X: 3213584, Y: 3463821, B: &Beacon{X: 3102959, Y: 3443573}},
	{X: 37652, Y: 3969055, B: &Beacon{X: -615866, Y: 3091738}},
	{X: 1804153, Y: 1170987, B: &Beacon{X: 1207332, Y: 429175}},
}
