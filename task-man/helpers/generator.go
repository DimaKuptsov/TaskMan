package helpers

import "math/rand"

func GenerateRandomString(length int, chars string) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = chars[GenerateIntBetween(0, len(chars))]
	}
	return string(bytes)
}

func GenerateIntBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}
