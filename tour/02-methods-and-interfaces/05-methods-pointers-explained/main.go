package main

import (
	"fmt"
	"math"
)

// Vertex struct of X, Y float64
type Vertex struct {
	X, Y float64
}

// Abs calculate the vertex length by using dot product against itself
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale the vertex by certain factor
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
