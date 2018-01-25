package secret

import (
	"strconv"
)

var m = map[int]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake implements secret handshake protocol.
func Handshake(input uint) []string {
	output := []string{}
	b := strconv.FormatInt(int64(input), 2)
	for i := 0; i < len(b) && i < 4; i++ {
		val := 1 << uint(i)
		if cur := int(input) & val; cur > 0 {
			output = append(output, m[int(val)])
		}
	}
	if int(input) > 16 {
		end := len(output) - 1
		for i := 0; i < len(output)/2; i++ {
			output[i], output[end-i] = output[end-i], output[i]
		}
	}

	return output
}
