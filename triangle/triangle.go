package triangle

import "math"

// Kind type of triangle.
type Kind int

const (
	NaT Kind = 0 // Nat not a triangle.
	Equ Kind = 1 // Equ equilateral.
	Iso Kind = 2 // Iso isosceles.
	Sca Kind = 3 // Sca scalene.
)

// KindFromSides returns type of triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a+b < c || b+c < a || c+a < b || a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a != b && b != c && a != c {
		k = Sca
	} else {
		k = Iso
	}
	return k
}
