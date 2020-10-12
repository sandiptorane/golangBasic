package main

import (
	"Module1/employee"
	"fmt"
)
func main(){
	details := employee.Details{
			Firstname : "sandip",
			Lastname : "torane",
			Salary : 15000,
	}
	fmt.Println("firstname :",details.Firstname," lastname :",details.Lastname," Salary:",details.Salary)
}
