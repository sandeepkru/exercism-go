package diffsquares

// SquareOfSums returns square of sum of first n numbers.
func SquareOfSums(n int) int {
	// calculate sum of numbers first - N * (N +1) / 2
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns sum of first n numbers squares.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns diff between SquareOfSums - SumOfSquares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
