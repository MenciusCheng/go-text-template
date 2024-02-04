package mock

import (
	"math/rand"
)

func generateRandomString(length int) string {
	//rand.Seed(time.Now().UnixNano())

	const charset = "0123456789"
	str := make([]byte, length)
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}

	return string(str)
}
