package main

import "fmt"

// Vertex contains Lat, Long float64
type Vertex struct {
	Lat, Long float64
}

// The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be
// added.
var m map[string]Vertex

func main() {
	// The `make` function returns a map of the given type, initialized and
	// ready for use.
	// If `:=` statement is used, declaration of m above is not needed.
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
