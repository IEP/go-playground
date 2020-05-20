package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk will traverse the tree leafs
func Walk(t *tree.Tree, ch chan int) {
	goWalk(t, ch)
	close(ch)
}

func goWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		goWalk(t.Left, ch)
		ch <- t.Value
		goWalk(t.Right, ch)
	}
}

// Same will compare tree t1 and t2 and return whether both same or different
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}

		// Both channel already closed indicates that two tree compared is
		// identical since it already passes above condition
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	// Check walk
	t := tree.New(1)
	ch := make(chan int)
	go Walk(t, ch)
	for v := range ch {
		fmt.Println(v)
	}

	// Compare 1 vs 1
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))

	// Compare 1 vs 2
	t1 = tree.New(1)
	t2 = tree.New(2)
	fmt.Println(Same(t1, t2))

	// Comapre 2 vs 1
	t1 = tree.New(2)
	t2 = tree.New(1)
	fmt.Println(Same(t1, t2))
}
