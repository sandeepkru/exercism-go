// Package isogram validates a string isogram or not.
package isogram

import "strings"

const alpha = "abcdefghijklmnopqrstuvwxyz"

// IsIsogram function takes string as input and returns true or false.
func IsIsogram(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	m := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if m[s[i]] == true && strings.Contains(alpha, string(s[i])) {
			return false
		}

		m[s[i]] = true
	}

	return true
}
