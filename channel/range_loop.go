package main

import "fmt"

func operations(channel chan int){
	for i:=0; i<5;i++{
		channel <-i
	}
	close(channel)
}

func main()  {
	ch := make(chan int)
	go operations(ch)
	for value:=range ch{
		fmt.Println("Received ",value)
	}
}

