package util

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	numbers  = "1234567890"
	specials = "!@#$%&*"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer in the interval [min,max).
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// randomString generates random string with the given length.
func randomString(length int) string {
	l := len(alphabet)

	bs := make([]byte, length)
	for i := 0; i < length; i++ {
		bs[i] = alphabet[rand.Intn(l)]
	}

	return string(bs)
}

// randomStringNumbers generates random string of numbers with the given length.
func randomStringNumbers(length int) string {
	l := len(numbers)

	bs := make([]byte, length)
	for i := 0; i < length; i++ {
		bs[i] = numbers[rand.Intn(l)]
	}

	return string(bs)
}

// randomStringSpecial generates random string of special symbols with the given length.
func randomStringSpecial(length int) string {
	l := len(specials)

	bs := make([]byte, length)
	for i := 0; i < length; i++ {
		bs[i] = specials[rand.Intn(l)]
	}

	return string(bs)
}

// RandomShortString generates random short string (4 chars).
func RandomShortString() string {
	return randomString(4)
}

// RandomString generates random string (8 chars).
func RandomString() string {
	return randomString(8)
}

// RandomLongString generates random string (16 chars).
func RandomLongString() string {
	return randomString(16)
}

// RandomEmail generates random valid email address.
func RandomEmail() string {
	return fmt.Sprintf("%s.%s@email.com", randomString(6), randomString(10))
}

// RandomPhoneNumber generates random valid phone number.
func RandomPhoneNumber() string {
	return randomStringNumbers(9)
}

// RandomPassword generates safe password.
func RandomPassword() string {
	a := randomString(10)
	b := randomStringNumbers(4)
	c := randomStringSpecial(2)

	return a + b + c
}
