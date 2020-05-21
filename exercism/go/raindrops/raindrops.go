// Package raindrops contain single function Convert that will convert a number
// into raindrop sounds.
package raindrops

import (
	"fmt"
	"strconv"
)

// Convert will convert number into raindrop sounds.
func Convert(num int) string {
	sound := ""

	// If num has 3 as a factor, add "Pling" sound.
	if num%3 == 0 {
		sound = fmt.Sprintf("%s%s", sound, "Pling")
	}

	// If num has 5 as a factor, add "Plang" sound.
	if num%5 == 0 {
		sound = fmt.Sprintf("%s%s", sound, "Plang")
	}

	// If num has 7 as a factor, add "Plong" sound.
	if num%7 == 0 {
		sound = fmt.Sprintf("%s%s", sound, "Plong")
	}

	// If num does not have any 3, 5, 7 as a factor the sound will be the num
	// itself.
	if sound == "" {
		sound = strconv.Itoa(num)
	}

	return sound
}
