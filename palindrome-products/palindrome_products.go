package palindrome

import (
	"errors"
	"math"
	"strconv"
)

// Product reprasents simple palindromic product and it's pairs.
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Products returns all palindromic products in a given range.
func Products(fmin, fmax int) (pmin, pmax Product, e error) {

	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}

	var palindromeExist bool
	pmin.Product = math.MaxInt32
	// Bruteforce, there should be better of doing this.
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			curProd := i * j
			if isPalindrome(curProd) {
				palindromeExist = true
				if pmin.Product > curProd {
					pmin = Product{curProd, [][2]int{{i, j}}}
				} else if pmax.Product < curProd {
					pmax = Product{curProd, [][2]int{{i, j}}}
				} else if pmin.Product == curProd {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				} else if pmax.Product == curProd {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}
			}
		}
	}
	if !palindromeExist {
		return pmin, pmax, errors.New("no palindromes")
	}

	return pmin, pmax, e
}

func isPalindrome(prod int) bool {
	s := strconv.Itoa(prod)
	palindrome := true
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			palindrome = false
			break
		}
	}

	return palindrome
}

// go test -run ^NOTHING -bench 'BenchmarkPalindromeProducts'
// goos: darwin
// goarch: amd64
// BenchmarkPalindromeProducts-8   	     100	  17784652 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/palindrome-products	1.806s
