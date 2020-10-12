package main

import "fmt"

func assert(i interface{}){
	s ,ok := i.(int)
	fmt.Println(s,ok)
}
func main(){
	value := 55
	assert(value)
	str := "hello world"
	assert(str)
}
