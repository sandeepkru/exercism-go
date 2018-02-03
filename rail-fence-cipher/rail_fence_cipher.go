package railfence

import (
	"strings"
)

// Encode returns encoded string based on number of rails.
func Encode(s string, rails int) string {
	encoded := generateList(s, rails)
	return strings.Join(encoded, "")
}

// Decode returns decoded string based on number of rails
func Decode(s string, rails int) string {
	// There should be better way of calculating each string length instead of generating strings themselves.
	lists := generateList(s, rails)
	decoded := make([]string, rails)
	j := 0
	k := 0
	for i := 0; i < rails; i++ {
		k += len(lists[i])
		decoded[i] = s[j:k]
		j += len(lists[i])
	}
	up := false
	var result string
	i := 0
	j = 0
	count := 0
	// track index of each string.
	indices := make([]int, rails)
	for {
		if count == len(s) {
			break
		}
		i = indices[j]
		indices[j] = i + 1
		result += string(decoded[j][i])
		if up {
			j--
			if j == 0 {
				up = !up
			}
		} else {
			j++
			if j == rails-1 {
				up = !up
			}
		}
		count++
	}
	return result
}

func generateList(s string, rails int) []string {
	arrange := make([]string, rails)
	j := 0
	var up bool
	for i := 0; i < len(s); i++ {
		arrange[j] += string(s[i])
		if up {
			j--
			if j == 0 {
				up = !up
			}
		} else {
			j++
			if j == rails-1 {
				up = !up
			}
		}
	}
	return arrange
}
