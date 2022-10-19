package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Init(x, y float64) (self *Point) {
	p.x, p.y = x, y
	return p
}

func (p *Point) X() float64 { return p.x }

func (p *Point) Y() float64 { return p.y }

func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(
		math.Pow(p.x-other.X(), 2) + math.Pow(p.y-other.Y(), 2),
	)
}

func main() {
	a := (&Point{}).Init(1, 1)
	b := (&Point{}).Init(2, 2)

	fmt.Println(a.DistanceTo(b))
}
