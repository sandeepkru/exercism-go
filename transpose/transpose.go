package transpose

import (
	"math"
)

// Transpose returns transposed list of strings.
func Transpose(input []string) []string {
	maxlen := 0
	for _, cur := range input {
		maxlen = int(math.Max(float64(maxlen), float64(len(cur))))
	}
	result := make([]string, maxlen)

	for j := 0; j < len(input); j++ {
		cur := input[j]
		for i := 0; i < len(cur); i++ {
			if len(result[i]) < j {
				k := j - len(result[i])
				for k > 0 {
					result[i] += " "
					k--
				}
			}
			result[i] += string(cur[i])
		}
	}

	return result
}

// go test -run ^NOTHING -bench 'BenchmarkTranspose'
// goos: darwin
// goarch: amd64
// BenchmarkTranspose-8   	  100000	     12437 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/transpose	1.384s
