package main

import "fmt"

func main(){
	employeeSalary := make(map[string]int)
	employeeSalary["steve"]=12000
	employeeSalary["Sandip"]=14500
	employeeSalary["Jack"]=14000
	fmt.Println("EmployeeSalary map:",employeeSalary)

	employeeSalary2 := map[string]int {  //second way to intialize the map
        "steve": 12000,
        "jamie": 15000,
    }
    employeeSalary2["mike"] = 9000
    fmt.Println("employeeSalary2 map contents:", employeeSalary2)

}
