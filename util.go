package logoped

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

func CreateRandomString(runes []rune, length int) (string, error) {
	if len(runes) == 0 {
		return "", errors.New("Wrong rune list")
	}
	if length > 0 {
		rand.Seed(time.Now().UnixNano())
		var b strings.Builder
		for i := 0; i < length; i++ {
			b.WriteRune(runes[rand.Intn(len(runes))])
		}
		return b.String(), nil
	}
	return "", errors.New("Wrong length")
}
