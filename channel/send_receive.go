package main

import "fmt"

func hello(ch chan bool){
	fmt.Println("inside goroutine")
	ch <- true //wrtie data to done channel
}

func main(){
	ch := make(chan bool)
	go hello(ch)
	<- ch //receiving the data form ch
	fmt.Println("main terminated")
}