package main

import "fmt"

type person struct {
	firstname string
	lastname string
}
func (p person) hello(){
	fmt.Println(p.firstname,p.lastname)
}

func main(){
	p := person{
		firstname: "sandip",
		lastname: "torane",
	}
	defer p.hello()
    fmt.Println("Welcome")
}
