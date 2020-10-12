package main

import(
	"fmt"
	"unicode/utf8"

)

func main(){
	word1 := "Señor"
	fmt.Println("string1 =",word1)
	fmt.Println("length =",utf8.RuneCountInString(word1))
	fmt.Println("no. of bytes =",len(word1))

	word2 := "sandip"
	fmt.Println("string1 =",word2)
	fmt.Println("length =",utf8.RuneCountInString(word2))
	fmt.Println("no. of bytes =",len(word2))

}
