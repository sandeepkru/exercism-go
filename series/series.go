package series

// All returns all substrings of s of length n.
func All(n int, s string) []string {
	result := []string{}
	if n > len(s) {
		return result
	}

	start := 0
	end := n
	for end <= len(s) {
		result = append(result, s[start:end])
		start++
		end++
	}

	return result
}

// UnsafeFirst returns first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	var result string
	if n > len(s) {
		return result
	}

	result = s[0:n]
	return result
}

// First is safe implementation of getting first sub string.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		ok = false
		return
	}
	first = s[0:n]
	ok = true
	return
}
