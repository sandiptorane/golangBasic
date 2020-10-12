package main

import "fmt"

func main(){
	employeeSalary := map[string]int{
		"steve":12000,
		"sandip":15000,
		"jack": 14500,
	}
	employee := "jack"
	salary := employeeSalary[employee]  //if employee present return its salary else return 0
	fmt.Println("salary of ",employee," = ",salary)

}
