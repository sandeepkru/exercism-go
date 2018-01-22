// Package grains implements functionality of numbers of grains per square in chess board.
package grains

import "errors"

var m = make(map[int]uint64)

func buildMap() {
	var prev uint64
	for i := 1; i < 65; i++ {
		if i == 1 {
			m[1] = 1
			prev = 1
		} else {
			prev = prev * 2
			m[i] = prev
		}
	}
}

// Square returns number of grains on a given square block number.
func Square(num int) (uint64, error) {
	buildMap()
	val, ok := m[num]
	if ok {
		return val, nil
	}
	return 0, errors.New("Invalid square number")
}

// Total returns total number of grains in total chess board.
func Total() uint64 {
	var total uint64
	for _, val := range m {
		total += val
	}
	return total
}
