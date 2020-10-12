package main

import "fmt"

type employee struct {
	salary int
	location string
}
func main(){
	emp1 :=employee{
		salary: 12000,
		location: "mumbai",
	}
	emp2 := employee{
		salary: 12500,
		location: "Pune",
	}
	emp3 := employee{
		salary: 14550,
		location: "banglore",
	}
	employeeinfo := map[string]employee{
		"steve" : emp1,
		"sam"   : emp2,
		"jack"  : emp3,
	}
	fmt.Println("employee Info")
	for key,info := range employeeinfo{
		fmt.Println("name =",key,"Salary =",info.salary,"location =",info.location)
	}
}
