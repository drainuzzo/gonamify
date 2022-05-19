package randify

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

const vowels = "aeiou"

var SeededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func Namify(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[SeededRand.Intn(len(charset))])
	}
	return sb.String()
}
