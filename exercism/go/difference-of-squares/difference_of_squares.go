package diffsquares

// Here, math import is not required due to overhead that may introduced when
// doing type conversion from and to int <-> float64.
//
// Additionally, all of the function has O(1) time complexity thus no need to
// change it into loop that has O(n) time complexity.

// SquareOfSum computes the square of sum from num = 1 to n
//     (1 + 2 + 3 + ... + n)^2
// SquareOfSum(n) = (n^4 + 2n^3 + n^2)/4
func SquareOfSum(n int) int {
	return (n*n*n*n + 2*n*n*n + n*n) / 4
}

// SumOfSquares computes the sum of squares from num = 1 to n
//     1^2 + 2^2 + 3^2 + ... + n^2
// SumOfSquares(n) = (n * (n + 1) * (2*n + 1))/6
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference computes the difference of square of sum and sum of squares for a
// given n.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
