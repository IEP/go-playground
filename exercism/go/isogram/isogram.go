package isogram

import (
	"strings"
)

// IsIsogram will check whether given string is considered isogram or not
func IsIsogram(s string) bool {
	visited := make(map[byte]bool)
	// Convert given string to lower case since isogram is case insensitive
	s = strings.ToLower(s)
	for _, v := range []byte(s) {
		// check only a-z character
		if 'a' <= v && v <= 'z' {
			// if the character is already occurs before the current iteration
			// return false
			if visited[v] {
				return false
			}
			// flag current character
			visited[v] = true
		}
	}
	return true
}
