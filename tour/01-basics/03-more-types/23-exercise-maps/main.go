package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount will count each word occurence of the provided string.
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	// Separate into word.
	wordPool := strings.Fields(s)
	// Iterate each word and build count map.
	for _, word := range wordPool {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
