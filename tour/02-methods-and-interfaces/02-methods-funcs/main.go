package main

import (
	"fmt"
	"math"
)

// Vertex is a vertex
type Vertex struct {
	X, Y float64
}

// Abs compute vertex length
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
