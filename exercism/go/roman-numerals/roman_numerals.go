package romannumerals

import "errors"

// ToRomanNumeral converts arabic number into roman numerals. If the number is
// not exist in roman numerals then empty string and error will be returned
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", errors.New("the number is not existed in roman numerals")
	}
	s := ""
	for n > 0 {
		switch {
		// 1000 mark
		case n >= 1000:
			s += "M"
			n -= 1000
			break
		case n >= 900:
			s += "CM"
			n -= 900
			break
		// 500 mark
		case n >= 500:
			s += "D"
			n -= 500
			break
		case n >= 400:
			s += "CD"
			n -= 400
			break
		// 100 mark
		case n >= 100:
			s += "C"
			n -= 100
			break
		case n >= 90:
			s += "XC"
			n -= 90
			break
		// 50 mark
		case n >= 50:
			s += "L"
			n -= 50
			break
		case n >= 40:
			s += "XL"
			n -= 40
			break
		// 10 mark
		case n >= 10:
			s += "X"
			n -= 10
			break
		case n >= 9:
			s += "IX"
			n -= 9
			break
		// 5 mark
		case n >= 5:
			s += "V"
			n -= 5
			break
		case n >= 4:
			s += "IV"
			n -= 4
			break
		case n >= 1:
			s += "I"
			n--
			break
		default:
			break
		}
	}
	return s, nil
}
