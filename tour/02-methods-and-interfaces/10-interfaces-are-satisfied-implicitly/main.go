package main

import "fmt"

// I interface of M()
type I interface {
	M()
}

// T structure of S string
type T struct {
	S string
}

// M method means type T implements the interface I, but we don't need
// explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
