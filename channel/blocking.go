package main

import (
	"fmt"
	"time"
)

func hello(done chan bool){
	fmt.Println("hello goroutine going to sleep")
	time.Sleep(4*time.Second)  //hello will sleep for 4 seconds and blocking main for 4 seconds
	fmt.Println("hello goroutine awake and going to write to done")
	done <- true
}
func main(){
	done := make(chan bool)
	fmt.Println("main goroutine going to call hello  goroutine ")
	go hello(done)
	<-done
	fmt.Println("data received")

}
