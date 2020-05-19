package main

import (
	"fmt"
)

// Sqrt will compute the square-root of given number by using Newton's method
func Sqrt(x float64) float64 {
	// Initial value
	z := 1.0
	// 10 iterations to achieve certain level of accuracy. This part can be
	// modified into conditional loop that compares previous and current value
	// of z and terminate the loop if only the z value already met the
	// required value of accuracy.
	for i := 0; i < 10; i++ {
		// The Newton's method is formulated as x_{n+1} = x_n - f(x_n)/f'(x_n)
		// Thus sqrt(x) can be approached by finding z value that fulfill
		//     z = sqrt(x)
		//     z^2 = x
		//     z^2 - x = 0
		// with x set to be constant. Thus the equation to find the value of z
		// can be expressed as
		//     z_{n+1} = z_n - f(z_n)/f'(z_n)
		//     z_{n+1} = z_n - (z^2 - x)/(2 * z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
