package main

import "fmt"

func assert(i interface{}){
	s := i.(int) //get underlying  int value from i
	fmt.Println(s)
}
func main(){
	value := 55
	assert(value)
}
