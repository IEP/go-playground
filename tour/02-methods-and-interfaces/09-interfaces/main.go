package main

import (
	"fmt"
	"math"
)

// Abser interface
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v // Error
	fmt.Println(a.Abs())
}

// MyFloat contains float64
type MyFloat float64

// Abs calculate absolute value of given number receiver
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex is a structure of X, Y float64
type Vertex struct {
	X, Y float64
}

// Abs calculate Vertex length
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
