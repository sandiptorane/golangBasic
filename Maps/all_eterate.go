package main

import (
	"fmt"
)
func main(){
	employeeSalary := map[string]int{
		"Sandip": 14500,
		"steve" : 12000,
		"jack" :  10000,
	}
	fmt.Println("contents in map=")
	for key,value:= range employeeSalary{
		fmt.Println("name=",key," Salary= ",value)
	}
}

