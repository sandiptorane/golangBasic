package main

import "fmt"

func printChar(s string){
	fmt.Printf("Characters: ")
	runes := []rune(s)    // ñ creates problem thats why use rune
	for i :=0; i < len(runes); i++{
		fmt.Printf("%c",runes[i])
	}
}
func main(){
	name := "Hello World"
	fmt.Printf("name = %s\n",name)
	printChar(name)
	fmt.Println()
	name= "Señor"
	fmt.Printf("String: %s\n", name)
	printChar(name)

}
