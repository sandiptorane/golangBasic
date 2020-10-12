package main

import "fmt"

func main(){
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	emp :="steve"
	value,flag := employeeSalary[emp]
	if flag==true{
		fmt.Println("Name =",emp,"Salary =",value)
		return
	}
	fmt.Println("employee not found")
}
