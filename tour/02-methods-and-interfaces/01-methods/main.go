package main

import (
	"fmt"
	"math"
)

// Vertex contains struct of X, Y float64
type Vertex struct {
	X, Y float64
}

// Abs calculate the vector length by using dot product against itself
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
