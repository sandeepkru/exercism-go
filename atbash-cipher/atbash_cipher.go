package atbash

import (
	"log"
	"regexp"
	"strings"
)

// Atbash takes normal string and returns atbash ciper string.
func Atbash(s string) string {
	var result string
	s = strings.ToLower(s)
	reg, err := regexp.Compile("[^a-z0-9]+")

	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, "")
	chars := []rune(s)
	for i, r := range chars {
		if r < 'a' || r > 'z' {
			result += string(r)
		} else {
			displacement := 'z' - r
			cur := 'a' + displacement
			result += string(cur)
		}
		if (i+1) < len(s) && (i+1)%5 == 0 {
			result += " "
		}
	}
	return result
}

// go test -run ^NOTHING -bench 'BenchmarkAtbash'
// goos: darwin
// goarch: amd64
// BenchmarkAtbash-8   	   20000	     99535 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/atbash-cipher	3.004s
