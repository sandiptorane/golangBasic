package main

import "fmt"

func main(){
	employeeSalary := make(map[string]int)
	employeeSalary["steve"]=12500
	employeeSalary["sandip"]=14500
	employeeSalary["smith"]=14820

	empto_delete :="smith"
	fmt.Println("before delete map contains =",employeeSalary)
	delete(employeeSalary,empto_delete)
	fmt.Println("after deleting employee map contains =",employeeSalary)
}
