// Package hamming contain function Distance that used to compute hamming
// distance of two given DNA strand
package hamming

import "errors"

// Distance will compute hamming distance of two given DNA strand, a and b.
// If the length of a and b is dissimilar, Distance will return error.
func Distance(a, b string) (int, error) {
	// If a and b has dissimilar then raise error.
	if len(a) != len(b) {
		return 0, errors.New("a and b must have same length")
	}

	// Initiate the hamming distance
	distance := 0

	// Convert to rune to cover utf-8 string
	strandA := []rune(a)
	strandB := []rune(b)

	// Iterate two DNA strands and compare it by characters at the same index.
	for index := range strandA {
		if strandA[index] != strandB[index] {
			distance++
		}
	}

	return distance, nil
}
