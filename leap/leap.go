package leap

// IsLeapYear checks if an year is leap year or not.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
