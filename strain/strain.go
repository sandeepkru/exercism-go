package strain

// Ints is collection of integers.
type Ints []int

// Lists is collection of integer sclices;
type Lists [][]int

// Strings is collection of strings.
type Strings []string

// Keep implements keep functionality on Ints.
func (its Ints) Keep(pred func(int) bool) Ints {
	var result Ints
	for _, cur := range its {
		if pred(cur) {
			result = append(result, cur)
		}
	}

	return result
}

// Discard implements discard functionality on Ints.
func (its Ints) Discard(pred func(int) bool) Ints {
	var result Ints
	for _, cur := range its {
		if !pred(cur) {
			result = append(result, cur)
		}
	}

	return result
}

// Keep implements keep functionality on Lists.
func (lists Lists) Keep(pred func([]int) bool) Lists {
	var result Lists
	for _, cur := range lists {
		if pred(cur) {
			result = append(result, cur)
		}
	}

	return result
}

// Keep implements keep functionality on Strings.
func (slist Strings) Keep(pred func(string) bool) Strings {
	var result Strings
	for _, cur := range slist {
		if pred(cur) {
			result = append(result, cur)
		}
	}

	return result
}

// go test -run ^NOTHING -bench 'BenchmarkKeepInts|BenchmarkDiscardInts'
// goos: darwin
// goarch: amd64
// BenchmarkKeepInts-8      	 5000000	       328 ns/op
// BenchmarkDiscardInts-8   	 5000000	       333 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/strain	4.011s
