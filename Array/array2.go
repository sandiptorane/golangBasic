package main

import(
	"fmt"
)

func main(){
	arr := [...]int{11,22,33,44} //compiler determines the length
	fmt.Println(arr)

        arr2 :=arr
	arr2[0]=66
	fmt.Println("after copying and updating arr to arr2:",arr2)

}

