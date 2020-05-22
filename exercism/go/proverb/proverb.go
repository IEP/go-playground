package proverb

import "fmt"

// Proverb will generate "For Want of a Nail" for the given rhyme
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))
	for i := range rhyme {
		// At the last index, use the last phrase instead of for want ... phrase
		if i == len(rhyme)-1 {
			proverb[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
	}
	return proverb
}
