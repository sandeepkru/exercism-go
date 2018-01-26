package summultiples

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// SumMultiples returns sum of all multiples up to the given limit.
func SumMultiples(limit int, divisors ...int) int {
	var total int
	for i := 1; i < limit; i++ { // Hardway of doing it.
		counted := false
		for _, cur := range divisors {
			if i%cur == 0 && !counted {
				total += i
				counted = true
			}
		}
	}
	return total
	// for _, cur := range divisors { // Some bug in the code, didn't work.
	// 	upto := limit / cur
	// 	curtotal := 0
	// 	if limit%cur == 0 {
	// 		upto--
	// 	}
	// 	for i := 1; i <= upto; i++ {
	// 		curtotal += (cur * i)
	// 	}
	// 	total += curtotal
	// }

	// if len(divisors) >= 2 {
	// 	for i := 1; i < len(divisors); i++ {
	// 		lcm := lcm(divisors[i-1], divisors[i])
	// 		upto := limit / lcm
	// 		if limit%lcm == 0 {
	// 			upto--
	// 		}
	// 		subs := 0
	// 		for i := 1; i <= upto; i++ {
	// 			subs += (lcm * i)
	// 		}
	// 		total -= subs
	// 	}

	// }

	// return total
}
