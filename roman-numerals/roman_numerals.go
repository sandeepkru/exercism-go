package romannumerals

import "errors"

// https://www.geeksforgeeks.org/converting-decimal-number-lying-between-1-to-3999-to-roman-numerals/
var m = []string{"", "M", "MM", "MMM"}
var c = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var x = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var i = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

// ToRomanNumeral converts given arabic number to roman numeral.
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", errors.New("Can't convert to roman")
	}

	var result string
	result = m[n/1000]
	result += c[(n%1000)/100]
	result += x[(n%100)/10]
	result += i[n%10]
	return result, nil
}

// go test -run ^NOTHING -bench 'BenchmarkRomanNumerals'
// goos: darwin
// goarch: amd64
// BenchmarkRomanNumerals-8   	 1000000	      1176 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/roman-numerals	1.197s
