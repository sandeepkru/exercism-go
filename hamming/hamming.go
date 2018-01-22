// Package hamming implements functionality of calculating hamming distance.
package hamming

import "errors"

// Distance function takes two DNA strands and returns computed hamming distance between two.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Invalid Length")
	}

	hdistance := 0
	for i := range a {
		if a[i] != b[i] {
			hdistance++
		}
	}
	// for i := 0; i < len(a); i++ {
	// 	if a[i] != b[i] {
	// 		hdistance++
	// 	}
	// }

	return hdistance, nil
}
