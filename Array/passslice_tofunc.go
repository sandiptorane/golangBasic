package main

import "fmt"

func operation(number []int){
	for i:= range number{
		number[i] +=2	
	}
}

func main(){
	nums := []int{11,22,33}
	fmt.Println("Slice before  function call=",nums)
	operation(nums)
	fmt.Println("slice after function call=",nums)

}
