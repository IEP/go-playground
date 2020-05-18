package main

import (
	"fmt"
	"math"
)

func main() {
	// Initially the tour provide `math.pi` name, but it against the convention
	// of `exported names` that need to start with capital letter. Thus it need
	// to change `math.pi` into `math.Pi`.
	fmt.Println(math.Pi)
}
