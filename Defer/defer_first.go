package main

import "fmt"

func finished(){
	fmt.Println("finished function operations")
}

func largest(num []int){
	defer finished()
	fmt.Println("largest funtions starts")
	large := num[0]
	for _,value := range num{
		if value > large{
			large = value
		}
	}
	fmt.Println("largest number in ",num, "is",large)
}

func main(){
	number := []int{11,55,22,33,44,35}
	largest(number)

}