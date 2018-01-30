package sieve

// Sieve returns all prime numbers upto given number.
func Sieve(limit int) []int {
	var primes []int
	if limit < 2 {
		return primes
	}

	if limit == 2 {
		primes = append(primes, 2)
		return primes
	}

	// map to maintain indication of prime number or not.
	m := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		v, ok := m[i]
		if ok && v {
			continue
		} else {
			primes = append(primes, i)
			// mark all multiples now
			for j := 2; j*i <= limit; j++ {
				m[i*j] = true
			}
		}
	}
	return primes
}

// go test -run ^NOTHING -bench 'BenchmarkSieve'
// goos: darwin
// goarch: amd64
// BenchmarkSieve-8   	   10000	    147167 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/sieve	1.498s
