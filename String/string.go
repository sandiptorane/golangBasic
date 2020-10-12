package main

import "fmt"

func printbyChar(s string){
	fmt.Printf("Characters: ")
	for i :=0; i < len(s); i++{
		fmt.Printf("%c",s[i])
	}
}

func printbyBtye(s string){
	fmt.Printf("Bytes:")
	for i :=0; i< len(s);i++{
		fmt.Printf("%X",s[i])
	}
}

func main(){
	name := "Hello World"
	fmt.Printf("name = %s\n",name)
	printbyChar(name)
	fmt.Printf("\n")
	printbyBtye(name)
}
