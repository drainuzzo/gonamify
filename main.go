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
		name := rndString(10, charset)
		fmt.Println("Name", i, " :", checkCons(name))

		//fmt.Println("Name", i, " :", rndString(1, vowels))
	}

	//fmt.Println(isCons('c', cns))
	//fmt.Println(gotVowel("ccecgpdrr"))
	//fmt.Println(!gotVowel("cccgpdrr"))
	//str1 := "cciolo"
	//str2 := "ccciolo"
	//str3 := "ccccolo"
	//str4 := "ccccclo"
	//str5 := "cccccco"
	//str6 := "ccccccc"

	//for i := 1; i < 7; i++ {
	//fmt.Println(i, " nome=", str1, "lavorato=", checkCons(str1))
	//fmt.Println(i, " nome=", str2, "lavorato=", checkCons(str2))
	//fmt.Println(i, " nome=", str3, "lavorato=", checkCons(str3))
	//fmt.Println(i, " nome=", str4, "lavorato=", checkCons(str4))
	//fmt.Println(i, " nome=", str5, "lavorato=", checkCons(str5))
	//fmt.Println(i, " nome=", str6, "lavorato=", checkCons(str6))
	//}
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
	return strings.ContainsAny(s, string(vowels))

}

//TO_FIX
func checkCons(name string) string {
	var str string = ""
	//fmt.Println(" ******* START *******\n\n")
	//for i := 0; i < len(name)-2; i++ {

	//fmt.Println("\n string i i+1 i+2=:", string(name[0]), string(name[1]), string(name[2]))
	if len(name) > 3 {
		if !gotVowel(string(name[0])) && !gotVowel(string(name[1])) && !gotVowel(string(name[2])) { //is consonant
			//fmt.Println(!gotVowel(string(name[0])) && !gotVowel(string(name[1])) && !gotVowel(string(name[2])))
			//fmt.Println("str prima ", str)
			str = string(name[0]) + rndString(1, vowels) + string(name[2])
			//fmt.Println("str dopo ", str)
			//fmt.Println("\n***** FOUND! **** giro \n")
		} else {
			if !gotVowel(string(name[2])) && !gotVowel(string(name[3])) {
				str = string(name[0]) + string(name[1]) + rndString(1, vowels)
			} else {
				str = string(name[0]) + string(name[1]) + string(name[2])
			}

			//fmt.Println("\n ****** ELSE ******* \n")

		}

		return str + checkCons(name[3:])
	}
	//fmt.Println("len name: ", len(name), "str: ", str)
	//s[randify.SeededRand.Intn(len(s))]
	// if !gotVowel(string(str[len(str)-1])) && !gotVowel(string(str[len(str)-2])) {
	// 	name = name[:len(name)-1] + rndString(1, vowels)
	// }
	return name

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
