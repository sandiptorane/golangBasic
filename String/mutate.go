package main

import "fmt"

func mutate(str []rune)string{
	str[0]='a'
	return string(str)
}

func main(){
	s := "Hello"
	fmt.Println("before modify :",s)
	fmt.Println("after modify :",mutate([]rune(s)))
}
