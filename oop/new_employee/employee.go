package new_employee

import "fmt"

type employee struct{
	firstname string
	lastname string
	totalLeaves int
	leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  //act as connstructor
	e := employee {firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee)CalculateLeaves(){
	fmt.Println(e.firstname ,e.lastname," has leaves remaining =",(e.totalLeaves-e.leavesTaken))
}
