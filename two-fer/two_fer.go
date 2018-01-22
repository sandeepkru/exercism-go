// Package twofer implements Two-fer.
package twofer

// ShareWith defines simple functionality of 2-fer i.e. it accepts a name
// and returns "One for you, one for me" string pattern.
func ShareWith(s string) string {
	if s == "" {
		return "One for you, one for me."
	}
	return "One for " + s + ", one for me."
}
