package main

import "fmt"

func main(){
	arr := [...]int{1,2,3,4,5}
	fmt.Println("before function call: ",arr)
	changearray(arr)
	fmt.Println("after function call: ",arr)
}

func changearray(arr [5]int){
	arr[0]=66
	fmt.Println("inside function:",arr)
}
