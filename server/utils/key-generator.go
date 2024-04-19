package utils

import (
	"math/rand"
	"time"
)

func GenerateUniqueKey() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 12)

	for idx := range bytes {
		bytes[idx] = letters[r.Intn(len(letters))]
	}

	return string(bytes)
}
