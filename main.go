package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func main() {
	// work in progress....

	fmt.Println("1:", rand.Int())
	fmt.Println("2:", rand.Int())
	fmt.Println("3:", rand.Int())
	fmt.Println("4:", seededRand.Int())
}

//TODO
func rndVowel([]string) {
	return //randomize vowels after 2? consonants
}

// func rndCons([]string) {
// 	return //randomize consonants
// }

/* TOCHECK
package main
import (
"fmt"
"math/rand"
"strings"
"time"
)
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func randomString(n int) string {
sb := strings.Builder{}
sb.Grow(n)
for i := 0; i < n; i++ {
sb.WriteByte(charset[rand.Intn(len(charset))])
}
return sb.String()
}
func main() {
rand.Seed(time.Now().UnixNano())
fmt.Println(randomString(20))
}
*/
