package main

import "fmt"

type salaryCalculator interface {
	calculateSalary()
}

type leavecalulator interface {
	calculateLeave() int
}

type employee struct {
	firstname string
	lastname string
	basicpay int
	pf int
	totalleaves int
	leavestaken int
}

func (e employee) calculateSalary() {
	fmt.Println("name",e.lastname,"surname",e.lastname,"salary =",e.basicpay+ e.pf)
}

func (e employee) calculateLeave()int{
	return e.totalleaves-e.leavestaken
}
func main(){
	e := employee{
		firstname: "sandip",
		lastname: "torane",
		basicpay: 70000,
		pf : 700,
		totalleaves: 30,
		leavestaken: 3,
	}
	var s salaryCalculator
	s =  e
	s.calculateSalary()
	var l leavecalulator = e
	fmt.Println("leaves remaining",l.calculateLeave())
}