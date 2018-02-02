package allergies

import "math"

var allergiesMap = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// Allergies reutrn list of allergies based on given score.
func Allergies(score uint) []string {
	var allergies []string

	for score >= 1 {
		nearest := math.Floor(math.Log2(float64(score)))
		val := math.Pow(2, nearest)
		if nearest > 7 {
			score = score % uint(val)
			continue
		} else {
			allergies = append([]string{allergiesMap[uint(val)]}, allergies...)
			score = score % uint(val)
		}
	}
	return allergies
}

// AllergicTo returns if person is alergic to or not based on given score and substance.
func AllergicTo(score uint, s string) bool {

	allergies := Allergies(score)

	for _, a := range allergies {
		if a == s {
			return true
		}
	}

	return false
}

// go test -run ^NOTHING -bench 'BenchmarkAllergies|BenchmarkAllergicTo'
// goos: darwin
// goarch: amd64
// BenchmarkAllergies-8    	 3000000	       392 ns/op
// BenchmarkAllergicTo-8   	 1000000	      1068 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/allergies	2.684s
