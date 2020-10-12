package main

import "fmt"

type employee  struct{
	name string
	lastname  string
	age int
	salary float64
}

func main(){
	//creatinng struct specifying field name
	emp1 := employee{
		name : "sandip",
		lastname : "Torane",
		age : 23,
		salary: 15000.0,
	}
	//creatinng struct without specifying field name
	emp2 := employee{"dinesh","megale",24,14000.0}

	fmt.Println("first name:",emp1.name,"last name:",emp1.lastname,"age:",emp1.age,"salary:",emp1.salary)
	fmt.Println(emp2)

	//creating anonymous structs
	emp3 := struct{
		firstname string
		lastname string
		age int
		salary float64
	}{
		firstname:"dhiraj",
		lastname: "gurav",
		age: 23,
		salary: 15400,
	}

	fmt.Println("anonymous struct emp3 :",emp3)

	emp4 := employee{
		name: "steve",
		lastname: "jobs",

	}
	fmt.Println("emp4:",emp4)

}


