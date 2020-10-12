package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("inside goroutine")
}

func main(){
	go hello()
	time.Sleep(1*time.Second) //it will provide enough time to execute goroutine
	fmt.Println("Hello world")
}

