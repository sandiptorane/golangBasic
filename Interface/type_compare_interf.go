package main

import "fmt"

type describer interface {
	describe()
}

type person struct {
	firstname string
	lastname string
}

func (p person)describe(){
	fmt.Printf("firstname = %s lastname = %s",p.firstname,p.lastname)
}

func findType(i interface{}){
	switch v :=i.(type) {
	case describer:
		v.describe()
	default:
		fmt.Println("unknown type")

	}
}
func main(){
	findType("hello world")
	p := person{
		firstname : "sandip",
		lastname: "torane",
	}
	findType(p)
}
