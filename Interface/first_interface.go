package main

import "fmt"

//interface definition
type vowelFinder interface {
	findVowels() []rune
}

//
type myString string

//myString implements vowelFinder
func (ms myString) findVowels() []rune{
	var vowels []rune
	for _,v:= range ms{
		if v=='a' || v=='e' || v=='i' || v=='o' || v=='u'{
			vowels = append(vowels,v)
		}
	}
	return vowels
}

func main(){
	name := myString("Sandip Torane")
	var v vowelFinder
	v= name //possible since myytring implements voweFinder
	fmt.Printf("vowels in name: %c",v.findVowels())
}
