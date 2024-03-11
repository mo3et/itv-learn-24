package qcode

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y, Z float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func DistanceFunc() {
	p1 := Point{1, 2, 3}
	p2 := Point{4, 5, 6}
	fmt.Printf("Distance between p1 and p2: %f\n", distance(p1, p2))
}
