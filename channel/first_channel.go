package main

import "fmt"

func main(){
	var a chan int
	if a==nil{
		fmt.Printf("channel is nil \n")
		a = make(chan int)
		fmt.Printf("after type of channel %T\n",a)
	}
}
