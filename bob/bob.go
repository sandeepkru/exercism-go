// Package bob implements behavior of bob's responses on input.
package bob

import (
	"regexp"
	"strings"
)

// Hey takes input and returns Bob's response accordingly.
func Hey(remark string) string {
	response := "Fine. Be that way!"
	containsalpha := regexp.MustCompile(`[a-zA-Z]`).MatchString
	// trim all spaces
	remark = strings.TrimSpace(remark)
	if len(remark) > 0 {
		if strings.HasSuffix(remark, "?") {
			if strings.ToUpper(remark) == remark && containsalpha(remark) {
				response = "Calm down, I know what I'm doing!"
			} else {
				response = "Sure."
			}
		} else {
			if strings.ToUpper(remark) == remark && containsalpha(remark) {
				response = "Whoa, chill out!"
			} else {
				response = "Whatever."
			}
		}
	}

	return response
}
