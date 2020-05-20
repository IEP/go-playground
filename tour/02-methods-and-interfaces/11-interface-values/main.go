package main

import (
	"fmt"
	"math"
)

// I interface of M()
type I interface {
	M()
}

// T structure of S string
type T struct {
	S string
}

// M prints t.S
func (t *T) M() {
	fmt.Println(t.S)
}

// F type of float64
type F float64

// M prints f
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
