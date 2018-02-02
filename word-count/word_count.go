package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is word count map.
type Frequency map[string]int

// WordCount returns frequency of each word in given stream.
func WordCount(s string) Frequency {
	wc := make(Frequency)
	b, e := -1, -1

	s = strings.ToLower(s)
	for i, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if b == -1 {
				b = i
			}
			if e == -1 {
				e = i
			} else {
				e++
			}
		} else if char == '\'' && e != -1 {
			e++
		} else {
			if b != -1 && e != -1 {
				// if we get quote at end ignore it.
				if s[e] == '\'' {
					e--
				}
				word := s[b : e+1]
				wc[word]++
			}
			b, e = -1, -1
		}
	}
	if b != -1 && e != -1 {
		wc[s[b:e+1]]++
	}
	return wc
}

// Current approx might not work for few cases which are not mentioned here.
// go test -run ^NOTHING -bench 'BenchmarkWordCount'
// goos: darwin
// goarch: amd64
// BenchmarkWordCount-8   	  200000	      7001 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/word-count	1.489s
