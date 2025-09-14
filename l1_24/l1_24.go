package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(q *Point) float64 {
	xDist := p.x - q.x
	yDist := p.y - q.y
	return math.Sqrt(xDist*xDist + yDist*yDist)
}

func main() {
	a := NewPoint(3.0, -4.0)
	b := NewPoint(-6.0, 5.0)

	fmt.Printf("distance: %.2f\n", a.Distance(b))
}
