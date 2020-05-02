package strings

import (
	"math/rand"
	"time"
	"unsafe"
)

var randomize = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Random returns random string from letters with length
func Random(length int, letters string) string {
	bytes := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for index, cache, remain := length-1, randomize.Int63(), letterIdxMax; index >= 0; {
		if remain == 0 {
			cache, remain = randomize.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			bytes[index] = letters[idx]
			index--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&bytes))
}
