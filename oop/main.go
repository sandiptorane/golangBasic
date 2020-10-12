package main

import "oop/employee"

func main(){
	e := employee.Employee{
		Firstname :"sandip",
		Lastname : "torane",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.CalculateLeaves()
}