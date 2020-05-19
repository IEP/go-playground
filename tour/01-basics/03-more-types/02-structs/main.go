package main

import "fmt"

// Vertex expressed here have similar structure as in C
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
