package perfect

import (
	"errors"
	"math"
)

// Classification is a type indicating number classification.
type Classification int

const (
	// ClassificationDeficient is Nicomachus's indication of deficient number.
	ClassificationDeficient Classification = iota
	// ClassificationPerfect is Nicomachus's indiccation of perfect number.
	ClassificationPerfect Classification = iota
	// ClassificationAbundant is Nicomachus's indication of abundant number.
	ClassificationAbundant Classification = iota
)

// ErrOnlyPositive indicates Classify accepts only numbers > 0.
var ErrOnlyPositive error

// Classify classifies number based on it's factors sum.
func Classify(num int64) (Classification, error) {

	if num <= 0 {
		ErrOnlyPositive = errors.New("Invalid input")
		return Classification(-1), ErrOnlyPositive
	}

	if num == 1 {
		return ClassificationDeficient, nil
	}

	factorsSum := getFactorsSum(num)
	if factorsSum == num {
		return ClassificationPerfect, nil
	} else if factorsSum < num {
		return ClassificationDeficient, nil
	} else {
		return ClassificationAbundant, nil
	}
}

func getFactorsSum(num int64) int64 {
	var factors []int64
	var i int64
	factors = append(factors, 1)
	for i = 2; i < int64(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			if num/i == i {
				factors = append(factors, i)
			} else {
				factors = append(factors, i)
				factors = append(factors, num/i)
			}
		}
	}

	var sum int64
	for _, val := range factors {
		sum += val
	}

	return sum
}

// go test -run ^NOTHING -bench 'BenchmarkClassify'
// goos: darwin
// goarch: amd64
// BenchmarkClassify-8   	   10000	    185169 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/perfect-numbers	1.879s
