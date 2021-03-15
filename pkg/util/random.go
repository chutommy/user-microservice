package util

import (
	"math/rand"
	"strings"
	"time"
)

var (
	alphabet    = "abcdefghijklmnopqrstuvwxyz"
	alphabetCap = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers     = "0123456789"
	special     = "?!@#$%&*"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random number in range of [from,to].
func RandomInt(from, to int) int {
	return rand.Intn(to-from+1) + from
}

// RandomString returns a random string with the given length.
func RandomString(length int) string {
	var sb strings.Builder
	for l := 0; l < length; l++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}

	return sb.String()
}

// RandomEmail returns a randomly generated valid email string.
func RandomEmail() string {
	var sb strings.Builder

	sb.WriteString(RandomString(4))
	sb.WriteString(".")
	sb.WriteString(RandomString(6))
	sb.WriteString("@")
	sb.WriteString(RandomString(5))
	sb.WriteString(".")
	sb.WriteString(RandomString(2))

	return sb.String()
}

// RandomEmail returns a random phone number.
func RandomPhoneNumber() string {
	var sb strings.Builder

	for n := 0; n < 9; n++ {
		sb.WriteByte(numbers[rand.Intn(len(numbers))])
	}

	return sb.String()
}

// RandomPassword returns a randomly generated basic password.
func RandomPassword() string {
	return RandomString(8)
}

// RandomName returns a random name
func RandomName() string {
	return strings.Title(RandomString(4))
}
