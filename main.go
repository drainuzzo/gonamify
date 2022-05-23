package main

import (
	"fmt"
	"strings"

	"github.com/drainuzzo/gonamify/randify"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

const vowels = "aeiou"

//var cns = getCons(charset, vowels)

//var cns string = strings.ReplaceAll(charset, vowels, "") //consonants

//var seededRand *rand.Rand = rand.New(
//	rand.NewSource(time.Now().UnixNano()))

func main() {
	// work in progress....

	//fmt.Println("1:", rand.Int())
	//fmt.Println("2:", rand.Int())
	//fmt.Println("3:", rand.Int())
	//fmt.Println("4:", randify.SeededRand.Int())
	//var example string = "ciccciopelliccio"
	// for i, r := range example {
	// 	fmt.Println(i, r, string(r))
	// }
	//rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 8; i++ {

		fmt.Println("Name", i, " :", rndString(10, charset))
		fmt.Println("Name", i, " :", rndString(1, vowels))
	}

	//fmt.Println(isCons('c', cns))
	fmt.Println(gotVowel("ccecgpdrr"))
	fmt.Println(!gotVowel("cccgpdrr"))
	//fmt.Println(cns)

}

func rndString(n int, s string) string {
	sb := strings.Builder{} //buffer
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(s[randify.SeededRand.Intn(len(s))])
	}
	return sb.String()
}

func gotVowel(s string) bool {
	if strings.ContainsAny(s, string(vowels)) {
		return true
	}
	return false
}

//TO FIX
/* func getCons(c string, v string) string {
	var cns string
	for i := range c {
		cns = strings.ReplaceAll(string(c[i]), string(v[i]), "")
		fmt.Println(cns)
	}
	return cns
} */

/* func fixCons(name string) string {

	for i:=0; i<len(name); i++ {
		for j,_ := range cns {
			if name[i]==cns[j] name[i+1]==cns[j] {
				name[i+1]=rndString(1, vowels)
			}

		if (strings.ContainsAny(cns,name[i])) &&
			(strings.Contains(cs[i+1],cns)) {
				cs[i+1] = rndString(1, vowels)
		}
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
