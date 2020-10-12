package main

import  "fmt"

func main(){
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	name :=string(byteSlice)
	fmt.Println("name=",name)
	byteSlice =[]byte{67, 97, 102, 195, 169}
	name =string(byteSlice)
	fmt.Println("name=",name)
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println("rune to string =",str)
}