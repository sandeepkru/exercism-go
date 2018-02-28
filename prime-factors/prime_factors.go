package prime

import "math"

// Factors return slice of prime factors for given number.
func Factors(input int64) []int64 {
	factors := []int64{}

	for input%2 == 0 {
		factors = append(factors, int64(2))
		input /= 2
	}

	var i int64
	i = 3
	for ; i <= int64(math.Sqrt(float64(input))); i += 2 {
		for input%i == 0 {
			factors = append(factors, i)
			input /= i
		}
	}

	if input > 2 {
		factors = append(factors, input)
	}

	return factors
}

// go test -run ^NOTHING -bench 'BenchmarkPrimeFactors'
// goos: darwin
// goarch: amd64
// BenchmarkPrimeFactors-8   	   30000	     49111 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/prime-factors	1.978s
