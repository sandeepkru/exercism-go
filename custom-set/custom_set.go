package stringset

import (
	"reflect"
	"strings"
)

// Set is representation of string set.
type Set struct {
	set map[string]bool
}

// New returns new stringSet
func New() Set {
	newSet := Set{
		set: make(map[string]bool),
	}
	return newSet
}

// NewFromSlice returns Set from string slice.
func NewFromSlice(input []string) Set {
	newSet := New()
	for _, cur := range input {
		newSet.Add(cur)
	}

	return newSet
}

// IsEmpty checks if string set is empty or not.
func (s Set) IsEmpty() bool {
	return len(s.set) == 0
}

// Has checks if given string exists in set or not.
func (s Set) Has(input string) bool {
	return s.set[input]
}

// Add inserts into string Set.
func (s Set) Add(input string) {
	s.set[input] = true
}

// String returns string representation of string set.
func (s Set) String() string {
	var result []string
	for k := range s.set {
		result = append(result, "\""+k+"\"")
	}
	content := strings.Join(result, ", ")
	return "{" + content + "}"
}

// Subset checks if one set is subset of other
func Subset(s1, s2 Set) bool {
	for k := range s1.set {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

// Equal compares if two sets are equal or not.
func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1.set, s2.set)
}

// Disjoint evaluates disjoint equality of two sets.
func Disjoint(s1, s2 Set) bool {
	if len(s2.set) > len(s1.set) {
		s1, s2 = s2, s1
	}
	for k := range s2.set {
		if s1.Has(k) {
			return false
		}
	}

	return true
}

// Intersection returns new set as intersection of two sets.
func Intersection(s1, s2 Set) Set {
	// Make always s1 as bigger set
	if len(s2.set) > len(s1.set) {
		s1, s2 = s2, s1
	}
	newSet := New()
	for k := range s2.set {
		if s1.Has(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

// Union returns union of two sets.
func Union(s1, s2 Set) Set {
	newSet := New()
	for k := range s1.set {
		newSet.Add(k)
	}
	for k := range s2.set {
		newSet.Add(k)
	}

	return newSet
}

// Difference performs A - B set operation.
func Difference(s1, s2 Set) Set {
	newSet := New()
	for k := range s1.set {
		if s2.Has(k) {
			continue
		}
		newSet.Add(k)
	}

	return newSet
}

// go test -run ^NOTHING -bench 'BenchmarkNewFromSlice1e1|BenchmarkNewFromSlice1e2|BenchmarkNewFromSlice1e3|BenchmarkNewFromSlice1e4'
// goos: darwin
// goarch: amd64
// BenchmarkNewFromSlice1e1-8   	 3000000	       406 ns/op
// BenchmarkNewFromSlice1e2-8   	  200000	     10664 ns/op
// BenchmarkNewFromSlice1e3-8   	   20000	     92536 ns/op
// BenchmarkNewFromSlice1e4-8   	    2000	    857888 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/custom-set	8.516s
