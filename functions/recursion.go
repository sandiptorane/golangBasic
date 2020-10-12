package main

import "fmt"

func fact(n int) int{
	if n==0{
		return 1
	}
	return n*fact(n-1)

}

func main(){
	 var num int =7
	 fmt.Printf("factorial of %d = %d",num,fact(num))
}
