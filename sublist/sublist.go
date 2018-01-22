package sublist

import (
	"reflect"
)

// Relation is simple string representation of slices relation.
type Relation string

func isSublist(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}
	var match bool
	var prevMatch int
	var matchCount int
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		//fmt.Printf("i %d, j %d\n", i, j)
		if l1[i] == l2[j] {
			if !match {
				match = true
				prevMatch = j
			}
			i++
			j++
			matchCount++
		} else if l1[i] > l2[j] {
			if match {
				// reset i
				i = 0
				match = false
				j = prevMatch + 1
				matchCount = 0
			} else {
				j++
			}
		} else {
			if prevMatch != 0 {
				j++
			} else { // increment only i if we have never set prevMatch otherwise increment j.
				i++
			}
		}
	}

	return matchCount == len(l1)
}

// Sublist checks if lists are subset of each other or not.
func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 && len(l2) == 0 {
		return "equal"
	}
	if len(l1) == len(l2) {
		if reflect.DeepEqual(l1, l2) {
			return "equal"
		}
		return "unequal"
	}

	if len(l2) > len(l1) {
		if isSublist(l1, l2) {
			return "sublist"
		}
		return "unequal"
	}
	if isSublist(l2, l1) {
		return "superlist"
	}
	return "unequal"
}
