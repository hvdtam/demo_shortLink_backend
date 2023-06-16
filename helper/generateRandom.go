package helper

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandom(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var result strings.Builder
	for i := 0; i < length; i++ {
		n := random.Intn(36)
		if n < 26 {
			result.WriteString(string(rune('a' + n)))
		} else {
			result.WriteString(strconv.Itoa(n - 26))
		}
	}
	return result.String()
}
