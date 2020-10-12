package main

import "fmt"

func evaluate(a int){
	fmt.Println("value of a in defered function",a)
}

func main(){
	a := 5
	defer evaluate(a)
	a =10
	fmt.Println("value of a before defer call",a)
}