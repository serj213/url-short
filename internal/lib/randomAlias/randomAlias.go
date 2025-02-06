package randomalias

import (
	"math/rand"
	"time"
)

var RandomAlias = func (length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]rune, length)

	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}
	
	return string(result)
}