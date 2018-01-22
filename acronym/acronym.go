// Package acronym implements shortening of strings to their acronyms.
package acronym

import "strings"

// Abbreviate takes long string and returns it's short form.
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	words := strings.Split(s, " ")
	result := ""

	for _, word := range words {
		result += strings.ToUpper(string(word[0])) // extract just first letter
	}
	return result
}
