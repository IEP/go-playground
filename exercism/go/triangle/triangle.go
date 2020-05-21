// Package triangle contain a function KindFromSides that return triangle kind
// given length of three sides.
package triangle

import "math"

// Kind store constant value of triangle
type Kind int

const (
	// NaT = not a triangle
	NaT Kind = iota
	// Equ = equilateral triangle; has three same length
	Equ
	// Iso = isosceles triangle; has two same length
	Iso
	// Sca = scalene triangle
	Sca
)

// KindFromSides return the kind of triangle given length of three sides.
func KindFromSides(a, b, c float64) Kind {
	// Basic triangle inequality and infinity check
	if a+b >= c && b+c >= a && a+c >= b && !math.IsInf(a, 0) &&
		!math.IsInf(b, 0) && !math.IsInf(c, 0) {
		switch {
		// Negative length check
		case a <= 0 || b <= 0 || c <= 0:
			return NaT
		// Equilateral check
		case a == b && b == c:
			return Equ
		// Isosceles check
		case (a == b && a != c) || (b == c && b != a) || (a == c && a != b):
			return Iso
		// Default; return Scalene
		default:
			return Sca
		}
	}
	return NaT
}
