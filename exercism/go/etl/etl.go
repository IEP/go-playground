package etl

import "strings"

// Transform performs ETL from legacy source format into the newer one.
func Transform(s map[int][]string) map[string]int {
	d := make(map[string]int)
	for score, letterArray := range s {
		for _, letter := range letterArray {
			// The key destination format has lower case, so it's necessary to
			// transform it first into lower case.
			d[strings.ToLower(letter)] = score
		}
	}
	return d
}
