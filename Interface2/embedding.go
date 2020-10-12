package main

import (
	"fmt"
)

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type employeeOperation interface {
	SalaryCalculator
	LeaveCalculator
}

type employ struct {
	firstname string
	lastname string
	basicpay int
	pf int
	totalleaves int
	leavestaken int
}

func (e employ)DisplaySalary(){
	fmt.Println("firstname ",e.firstname,"lastname",e.lastname,"salary =",(e.basicpay+e.pf))
}

func (e employ)CalculateLeavesLeft() int{
	return (e.totalleaves-e.leavestaken)
}

func main(){
	e := employ{
		firstname: "sandip",
		lastname: "torane",
		basicpay: 70000,
		pf : 700,
		totalleaves: 30,
		leavestaken: 3,
	}
	var empOp employeeOperation = e
	empOp.DisplaySalary()
	fmt.Println("remaining leaves =",empOp.CalculateLeavesLeft())
}
