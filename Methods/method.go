package main

import "fmt"

type employee struct{
	name string
	currency string
	salary float64
}

func (e employee) displayEmploy(){
	fmt.Println("employee:",e.name,"currency :",e.currency,"salary:",e.salary)
}

func main(){
	emp1 := employee{
		name: "sandip",
		currency: "Rupees",
		salary : 15000,
	}
	emp1.displayEmploy()
}
