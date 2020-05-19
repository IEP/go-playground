package main

// Range continued
//
// You can skip the index or value by assigning to `_`.

import "fmt"

func main() {
	pow := make([]int, 10)

	// This one is similar to for i, _ := range pow.
	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}

	// This one omit the first variable, retains only the second.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
