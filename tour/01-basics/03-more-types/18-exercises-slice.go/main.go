package main

import "golang.org/x/tour/pic"

// Pic implement Pic in go tour
func Pic(dx, dy int) [][]uint8 {
	// Pre-allocate
	s := make([][]uint8, dy)
	for y := range s {
		s[y] = make([]uint8, dx)
	}
	// Process: using x ^ y
	for y := range s {
		for x := range s[y] {
			s[y][x] = uint8(x ^ y)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
