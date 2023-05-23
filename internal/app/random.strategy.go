package app

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	UrlSize       = 7
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

type RandomNumbersStrategy struct {
}

func (r *RandomNumbersStrategy) shorten(url string) (string, error) {
	url, err := getRandomString(UrlSize)
	return url, err
}

func getRandomString(size int) (string, error) {
	sb := strings.Builder{}
	sb.Grow(size)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := size-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String(), nil
}

func NewRandomShortener() *RandomNumbersStrategy {
	return &RandomNumbersStrategy{}
}
