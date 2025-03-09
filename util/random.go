package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabets = "abcdefghijklmnopqrstuvwxyx"

// RandomInt generates a random int b/w max and min
func RandomInt(max, min int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string of n characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}