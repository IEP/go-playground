package main

import (
	"fmt"
	"math"
)

// sqrt will compute the square-root of given number, if the number is negative
// the function will call itself with complement of the number
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
