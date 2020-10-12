package main

import "fmt"

//interface definition
type salaryCaclulator interface {
	calculateSalary() int
}

type  permanent struct{
	empId int
	basicPay int
	pf int
}

type contract struct{
	empId int
	basicPay int
}

func (p permanent) calculateSalary() int{
	return p.basicPay+p.pf
}

func (c contract) calculateSalary() int{
	return c.basicPay
}

//sum the each and every employee's salary and display total
func totalexpenses(s []salaryCaclulator) {
	expense := 0
	for _,emp := range s{
		expense = expense + emp.calculateSalary()
	}
	fmt.Println("total expense to company for employee salaries:",expense)
}
func main(){

	perm1 := permanent{
		empId: 1,
		basicPay: 15000,
		pf: 150,
	}
	perm2 := permanent{
		empId: 2,
		basicPay: 12000,
		pf: 120,
	}
	cont1 := contract{
		empId: 3,
		basicPay: 10000,
	}
	cont2 := contract{
		empId: 4,
		basicPay: 11000,
	}

	employee := []salaryCaclulator{perm1,perm2,cont1,cont2}
	totalexpenses(employee)


}
