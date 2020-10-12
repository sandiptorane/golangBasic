package main

import "fmt"

func modify(val *int){
	*val =200
}

func main(){
	b :=100
	a := &b
	fmt.Println("before function call b= ",b)
	modify(a)
	fmt.Println("after function call b= ",b)
}
