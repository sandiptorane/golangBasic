package main

import  "fmt"

func senddata(sendch chan<- int){ //here is unidirectional channel it only send data
	sendch<-10
}

func main(){
    channel := make(chan int)  //bidirectional channel
    go senddata(channel)
    fmt.Println(<-channel)
}
