package  employee

import "fmt"

type Employee struct{
	Firstname string
	Lastname string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee)CalculateLeaves(){
	fmt.Println(e.Firstname ,e.Lastname," has leaves remaining =",(e.TotalLeaves-e.LeavesTaken))
}