package main

import "fmt"

func hello(){
	fmt.Println("inside goroutine")
}

func main(){
	go hello() // did not print anything beacause immediatly goroutine returns and go to next line
	fmt.Println("Hello world")
}
