package main

import "fmt"

func main(){
	var arr [3]int //int array
	arr[0] =11
	arr[1] = 22
	arr[2] = 33
	fmt.Println("Array=",arr)

	arr2 := [5]int{11,22,33,44,55}
	fmt.Println("Array2=",arr2)

}
