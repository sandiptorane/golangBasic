package main

import "fmt"

func main(){
	ch := make(chan string,3)
	ch <- "sandip"
	ch <- "sam"
	ch <- "jack"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
