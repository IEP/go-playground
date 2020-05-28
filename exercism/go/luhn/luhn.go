package luhn

import (
	"strings"
	"unicode"
)

// Valid check whether given number is valid per the Luhn formula
func Valid(nums string) bool {
	// strip space
	nums = strings.ReplaceAll(nums, " ", "")
	// as per the problem statement, string of length 1 or less is invalid
	if len(nums) <= 1 {
		return false
	}

	// reverse the number for easier iteration
	nums = reverseString(nums)

	sum := 0

	for i, num := range nums {
		// if the current number is not a digit, then it has invalid character
		// thus it also invalid in term of Luhn formula
		if !unicode.IsDigit(num) {
			return false
		}

		// check whether the digit from the right is even or odd digit
		// both of the statement use int(num-'0') to convert from rune to int
		// to reduce overhead assuming only ASCII decimal digit is included
		switch i%2 != 0 { // even digit; since the array start from 0
		case true:
			// if the 2 * even digit number is greater or equal than 10,
			// substract 9 from that digit
			if digit := int(num-'0') * 2; digit >= 10 {
				sum += digit - 9
			} else {
				sum += digit
			}
		case false:
			sum += int(num - '0')
		}

	}

	// check whether the sum is modulo of 10
	if sum%10 == 0 {
		return true
	}
	return false
}

// reverseString will reverse given string (self-explanatory).
func reverseString(s string) string {
	bs := []rune(s)
	n := len(bs)
	for i := 0; i < n/2; i++ {
		bs[i], bs[n-1-i] = bs[n-1-i], bs[i]
	}
	return string(bs)
}
