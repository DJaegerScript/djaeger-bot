package utils

import (
	"math/rand"
)

func GenerateTag(length int) string {

	bytes := make([]byte, length)
	for i := 0; i < 5; i++ {
		bytes[i] = byte(randInt(65, 90))
	}

	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
