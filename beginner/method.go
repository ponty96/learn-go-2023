package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	xAxis float64
	yAxis float64
}

func (v *Vertex) Scale(f float64) {
	v.xAxis = v.xAxis * f
	v.yAxis = v.yAxis * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.xAxis*v.xAxis + v.yAxis*v.yAxis)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println("Before scaling: ", v, " Abs: ", v.Abs())
	v.Scale(5)
	fmt.Println("After scaling: ", v, " Abs: ", v.Abs())
}
