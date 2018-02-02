package anagram

import (
	"strings"
)

// Detect detects anagrams from provided list.
// this is based on map approach, we can simply sort and compare too.
func Detect(s string, candidates []string) []string {
	var tolower bool
	// if subject is in upper case don't convert to lower.
	if s != strings.ToUpper(s) {
		s = strings.ToLower(s)
		tolower = true
	}
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	var result []string
	for _, cur := range candidates {
		if len(cur) != len(s) {
			continue
		}
		temp := copyMap(m)
		curstr := cur
		if tolower {
			curstr = strings.ToLower(cur)
		}

		for _, c := range curstr {
			if val, ok := temp[c]; !ok {
				break
			} else {
				val--
				if val == 0 {
					delete(temp, c)
				} else {
					temp[c] = val
				}
			}
		}
		if len(temp) == 0 {
			result = append(result, cur)
		}
	}
	return result
}

func copyMap(m map[rune]int) map[rune]int {
	temp := make(map[rune]int)

	for k, v := range m {
		temp[k] = v
	}

	return temp
}

// go test -run ^NOTHING -bench 'BenchmarkDetectAnagrams'
// goos: darwin
// goarch: amd64
// BenchmarkDetectAnagrams-8   	  100000	     17389 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/anagram	1.930s
