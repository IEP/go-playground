package main

import "fmt"

// Vertex that consist of two fields: X and Y
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // The vertex X field modified
	fmt.Println(v.X)
}
