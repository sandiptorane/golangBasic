package main

import "fmt"

func operation(channel chan int){
	for i:=0; i<5;i++{
		channel <-i
	}
	close(channel)
}

func main(){
	channel := make(chan int)
	go operation(channel)
	for{
		value,ok := <-channel
		if(ok == false){
			break
		}
		fmt.Println("Received",value,ok)
	}

}