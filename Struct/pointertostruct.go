package main

import "fmt"

type employee  struct{
	name string
	lastname  string
	age int
	salary float64
}

func main(){
	emp5 := &employee{
		name:"sandip",
		lastname: "torane",
		age : 22,
		salary : 15240,
	}
	fmt.Println("firstname:",(*emp5).name)
	fmt.Println("lastname:",(emp5).lastname)  //we can use emp.name and (*emp).name both
	fmt.Println("age",emp5.age)
}

