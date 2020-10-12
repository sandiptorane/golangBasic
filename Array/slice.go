package main 

import "fmt"

func main(){
  	a := [5]int{21,22,23,24,25}
	var b []int = a[1:4] //create slice from a[1] to a[3]
	fmt.Println(b)

	c := []int{6,7,8}
	fmt.Println(c)
	for i := range b{
		b[i]++
	}
	fmt.Println("after modifying slice b it reflect in main arr a")
	fmt.Println("b=",b)
	fmt.Println("a=",a)
}
