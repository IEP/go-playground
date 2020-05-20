package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt type of float64
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

// Sqrt computes square root by using Newton's method
func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}
	z, prevZ := 1.0, 1000.0
	counter := 0
	for math.Abs(z-prevZ) > 1e-6 && counter < 20 {
		counter++
		prevZ, z = z, z-(z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
