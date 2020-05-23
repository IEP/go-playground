package scrabble

import (
	"strings"
)

// Score compute scrabble score of a given string.
func Score(s string) int {
	score := 0

	// Create score sheet to map characters and its corresponding score.
	scoreSheet := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}

	// Preprocess the given string: (1) lower the case and (2) split it to array
	// of character.
	// I prefer this way over creating histogram of character in sake of
	// readability. If the performance matters thus it better to refactor this
	// part into histogram of character and also modify the iteration.
	charArray := strings.Split(strings.ToLower(s), "")

	for scoreKey, scoreValue := range scoreSheet {
		for _, char := range charArray {
			// Check whether part of the given string (char) is a substring in
			// particular score sheet
			if strings.Contains(scoreKey, char) {
				score += scoreValue
			}
		}
	}
	return score
}
