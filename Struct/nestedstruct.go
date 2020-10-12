package main

import "fmt"

type Address struct{
	city string
	state string
}
type employee1 struct{
	firstname string
	lastname string
	addr Address
}

func main(){
	emp := employee1{
		firstname: "steve",
		lastname: "jobs",
		addr : Address{
			city: "pune",
			state: "maharashtra",
		},
	}
	fmt.Println("firstname =",emp.firstname)
	fmt.Println("lastname =",emp.lastname)
	fmt.Println("city =",emp.addr.city)
	fmt.Println("state =",emp.addr.state)

}
