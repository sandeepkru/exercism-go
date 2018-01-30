package lsproduct

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

// LargestSeriesProduct returns largest possible product from contiguous digits with given span.
func LargestSeriesProduct(s string, span int) (int64, error) {
	r := regexp.MustCompile("^[0-9]*$")
	if len(s) < span || span < 0 || !r.MatchString(s) {
		return 0, errors.New("Invalid input")
	}

	var product int64
	// this is a brute force method, there must be better DP solution for this.
	for i, j := 0, span; j <= len(s); i, j = i+1, j+1 {
		cur := s[i:j]
		var curProd int64
		curProd = 1
		for _, digit := range cur {
			num, err := strconv.Atoi(string(digit))
			if err != nil {
				return 0, errors.New("Invalid character in digits")
			}
			curProd *= int64(num)
		}
		product = int64(math.Max(float64(curProd), float64(product)))
	}

	return product, nil
}
