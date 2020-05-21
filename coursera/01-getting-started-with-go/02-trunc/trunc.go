package main

/* Write a program which prompts the user to enter a floating point number and
   prints the integer which is a truncated version of the floating point number
   that was entered. Truncation is the process of removing the digits to the
   right of the decimal place.
*/

import "fmt"

// trunc will truncate the number given
func trunc(num float64) int {
	return int(num)
}

func main() {
	var numFloat float64
	fmt.Scan(&numFloat)
	numInt := trunc(numFloat)
	fmt.Println(numInt)
}
