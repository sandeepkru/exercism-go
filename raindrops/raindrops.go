// Package raindrops provides string implementations based on factors.
package raindrops

import (
	"bytes"
	"strconv"
)

const (
	threefactor = "Pling"
	fivefactor  = "Plang"
	sevenfactor = "Plong"
)

// Convert function converts integer to string based on it's factors.
func Convert(num int) string {
	var buffer bytes.Buffer

	if num%3 == 0 {
		buffer.WriteString(threefactor)
	}

	if num%5 == 0 {
		buffer.WriteString(fivefactor)
	}

	if num%7 == 0 {
		buffer.WriteString(sevenfactor)
	}

	if buffer.Len() == 0 {
		buffer.WriteString(strconv.Itoa(num))
	}

	return buffer.String()
}
