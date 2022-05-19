package rand

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

const vowels = "aeiou"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func namify(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}
