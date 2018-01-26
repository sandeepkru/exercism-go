package pythagorean

// Triplet represents pythagorean triplet a, b, c where  a*a + b*b = c*c.
type Triplet [3]int

// Range returns all pythagorean triplets in given range.
func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

// Sum returns sum of pythagorean triplets.
func Sum(sum int) []Triplet {
	var triplets []Triplet
	for a := 1; a <= sum/2; a++ {
		for b := a; b <= sum/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}
