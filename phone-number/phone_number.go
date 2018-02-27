package phonenumber

import (
	"errors"
	"regexp"
)

// Number returns correctly formated phone number.
func Number(input string) (string, error) {
	if len(input) < 10 {
		return "", errors.New("Invalid input")
	}

	reg := regexp.MustCompile("[^0-9]+")
	input = reg.ReplaceAllString(input, "")

	if len(input) < 10 || len(input) > 11 {
		return "", errors.New("Invalid input")
	}

	if len(input) == 11 {
		if input[0] == '1' {
			input = input[1:]
		} else {
			return "", errors.New("Country code not starting with 1")
		}
	}

	// Validate exchange code
	if input[0] < '2' || input[0] > '9' || input[3] < '2' || input[3] > '9' {
		return "", errors.New("Invalid input")
	}

	return input, nil
}

// AreaCode returns area code of given phone number.
func AreaCode(input string) (string, error) {
	input, err := Number(input)
	if err != nil {
		return "", err
	}
	return input[:3], nil
}

// Format formats phone number as (xxx) yyy-zzzz
func Format(input string) (string, error) {
	input, err := Number(input)
	if err != nil {
		return "", err
	}

	formattedNum := "(" + input[:3] + ") " + input[3:6] + "-" + input[6:]
	return formattedNum, nil
}

// go test -run ^NOTHING -bench 'BenchmarkNumber|BenchmarkAreaCode|BenchmarkFormat'
// goos: darwin
// goarch: amd64
// BenchmarkNumber-8     	   20000	     86816 ns/op
// BenchmarkAreaCode-8   	   20000	     85764 ns/op
// BenchmarkFormat-8     	   20000	     86963 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/phone-number	7.827s
