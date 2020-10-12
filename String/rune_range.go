package main

import "fmt"
func charAndBytePosition(s string){
	for index,runes := range s{
		fmt.Printf("%c at position %d\n",runes,index)
	}
}

func main() {
	name := "Señor"
	charAndBytePosition(name)
}