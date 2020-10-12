package main

import "fmt"

func main(){
	name := "Naveen"
	fmt.Println("original string =",name)
	fmt.Printf("Reversed string = ")
	for _,v := range []rune(name){
		defer fmt.Printf("%c",v)
	}
}
