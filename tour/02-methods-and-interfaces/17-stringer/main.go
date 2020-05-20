package main

import "fmt"

// Person structure of Name and Age
type Person struct {
	Name string
	Age  int
}

// String receiver of Person type
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblerox", 9001}
	fmt.Println(a, z)
}
