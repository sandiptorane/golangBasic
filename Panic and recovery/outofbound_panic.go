package main

import "fmt"

func slicepanic(){
	num := []int{1,2,3,4}
	fmt.Println(num[5])
	fmt.Println("normally returned from a")
}
func main(){
	slicepanic()
	fmt.Println("successfully terminated from main")
}
