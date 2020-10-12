package main

import "fmt"

func concat(str1 string,str2 string )string{
	str3 := str1+" "+str2
	return str3
}

func main(){
	word1 := "Hello"
	word2 := "World"
	word3:= concat(word1,word2)
	fmt.Println("after concat word3 =",word3)
	result :=  fmt.Sprintf("%s %s",word1,word2) // add space betn concat
	fmt.Println("using Sprintf =",result)
}
