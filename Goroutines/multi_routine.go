package main

import (
	"fmt"
	"time"
)

func number(){
	for i:=0;i<5;i++{
		time.Sleep(200*time.Millisecond)
		fmt.Printf("%d ",i)
	}
}

func alphabets(){
	for i:='a';i<'f';i++{
		time.Sleep(350*time.Millisecond)
		fmt.Printf("%c ",i)
	}
}

func main(){
	go number()
	go alphabets()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("\nterminated")
}