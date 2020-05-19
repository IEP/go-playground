package main

import (
	"fmt"
	"math"
)

// MyFloat is a custom type that extends float64
type MyFloat float64

// Abs will compute the absolute value of f
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
