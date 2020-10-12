package main

import "fmt"

type Address struct{
	city string
	state string
}
type employee struct{
	firstname string
	lastname string
	Address
}

func main(){
	emp8 := employee{
		firstname: "jack",
		lastname:  "mick",
		Address: Address{ //promoted struct
			city:  "mumbai",
			state: "Maharashtra",
		},
	}
		fmt.Println("firstname =",emp8.firstname)
		fmt.Println("lastname =",emp8.lastname)
		fmt.Println("city =",emp8.city)          //promoted city
		fmt.Println("state =",emp8.state)

}
