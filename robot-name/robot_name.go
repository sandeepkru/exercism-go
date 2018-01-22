// Package robotname implements simple random robot name generation.
package robotname

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

// Got this from stackoverflow
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

)

// RandStringBytesMaskImprSrc RandomString generator.
func RandStringBytesMaskImprSrc() string {
	n := 2
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	n = 3
	c := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(numberBytes) {
			c[i] = numberBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b) + string(c)
}

// Robot is simple representation of robot.
type Robot struct {
	name string
}

// Name returns name of Robot.
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = RandStringBytesMaskImprSrc()
	}
	return r.name
}

// Reset resets robot to factory settings.
func (r *Robot) Reset() {
	r.name = ""
}
