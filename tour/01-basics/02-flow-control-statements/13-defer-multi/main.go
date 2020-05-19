package main

import "fmt"

func main() {
	fmt.Println("counting")

	// The deferred statement will be called in LIFO order after all the
	// expression other than the defer inside the function executed.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
