// Package luhn implements luhun algorithm to validate credit cards.
package luhn

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

// Valid validates given string is valid credit card or SIN based on luhn algorithm.
func Valid(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return false
	}

	skip := true

	var buffer bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			continue
		}
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
		if skip {
			skip = false
			buffer.WriteString(string(s[i]))
		} else {
			skip = true
			cur, _ := strconv.Atoi(string(s[i]))
			cur = cur * 2
			if cur > 9 {
				cur -= 9
			}
			buffer.WriteString(strconv.Itoa(cur))
		}
	}

	total := 0
	resultStr := buffer.String()
	for _, i := range resultStr {
		val, _ := strconv.Atoi(string(i))
		total += val
	}

	return total%10 == 0
}
