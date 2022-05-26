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
	//str1 := "cciolo"
	str2 := "ccciolo"
	// str3 := "ccccolo"
	// str4 := "ccccclo"
	// str5 := "cccccco"
	// str6 := "ccccccc"

	for i := 1; i < 7; i++ {
		//fmt.Println(i, " nome=", str1, "lavorato=", checkCons(str1))
		fmt.Println(i, " nome=", str2, "lavorato=", checkCons(str2))
		// fmt.Println(i, " nome=", str3, "lavorato=", checkCons(str3))
		// fmt.Println(i, " nome=", str4, "lavorato=", checkCons(str4))
		// fmt.Println(i, " nome=", str5, "lavorato=", checkCons(str5))
		// fmt.Println(i, " nome=", str6, "lavorato=", checkCons(str6))
	}
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

//TO_FIX
func checkCons(name string) string {
	var str string = ""
	fmt.Println("\n\n ******* START *******\n\n")
	for i := 0; i < len(name)-2; i++ {
		fmt.Println("\n string i i+1 i+2=", i, i+1, i+2, " :", string(name[i]), string(name[i+1]), string(name[i+2]))
		if !gotVowel(string(name[i])) && !gotVowel(string(name[i+1])) && !gotVowel(string(name[i+2])) { //is consonant
			str = str + string(name[i]) + rndString(1, vowels) + string(name[i+2]) + string(name[i+3:])
			fmt.Println("\nFOUND! itero i=", i, " str:", str)
		}
		return str + checkCons(name[i+1:len(name)-2])
		//? TODO
	}
	if str == "" {
		str += name
	}
	fmt.Println("\n ****** END *******\n\n")
	return str

}

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
