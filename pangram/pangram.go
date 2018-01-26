package pangram

import "strings"

// IsPangram validates given input is pangram or not.
func IsPangram(input string) bool {
	if input == "" {
		return false
	}

	charMap := make(map[rune]bool)
	input = strings.ToLower(input)
	for _, ch := range input {
		charMap[ch] = true
	}

	for i := 'a'; i <= 'z'; i++ {
		if _, ok := charMap[i]; !ok {
			return false
		}
	}
	return true
}
