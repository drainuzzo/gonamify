package main

import (
	"fmt"
	"strings"

	"github.com/drainuzzo/gonamify/randify"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

const vowels = "aeiou"

//var seededRand *rand.Rand = rand.New(
//	rand.NewSource(time.Now().UnixNano()))

func main() {
	// work in progress....

	//fmt.Println("1:", rand.Int())
	//fmt.Println("2:", rand.Int())
	//fmt.Println("3:", rand.Int())
	//fmt.Println("4:", randify.SeededRand.Int())
	var example string = "ciccciopelliccio"
	for i, r := range example {
		fmt.Println(i, r, string(r))
	}
	//rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 8; i++ {

		fmt.Println("Name", i, " :", rndString(10, charset))
		fmt.Println("Name", i, " :", rndString(1, vowels))
	}

}

func rndString(n int, s string) string {
	sb := strings.Builder{} //buffer
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(s[randify.SeededRand.Intn(len(s))])
	}
	return sb.String()
}

/* func fixCons(s *string) string {
	for i := 0; i < n; i++ {

	}
}  */

//TODO
// func rndVowel(string) {

// 	return //randomize vowels after 2? consonants
// }

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
